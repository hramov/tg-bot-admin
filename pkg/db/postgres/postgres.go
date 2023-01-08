package postgres

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/config"
	initDb "github.com/hramov/tg-bot-admin/pkg/db/postgres/init"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Postgres struct {
	db     *sqlx.DB
	logger *logging.Logger
}

var instance *Postgres

func Connect(cfg *config.Config, logger *logging.Logger) (*Postgres, error) {
	if instance != nil {
		return instance, nil
	}

	instance = &Postgres{}
	db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Storage.Username, cfg.Storage.Password, cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.Database, cfg.Storage.SslMode))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Storage.MaxOpenCons)
	db.SetMaxIdleConns(cfg.Storage.MaxIdleCons)
	db.SetConnMaxIdleTime(cfg.Storage.ConnMaxIdleTime)
	db.SetConnMaxLifetime(cfg.Storage.ConnMaxLifetime)

	err = initDb.Start(db)
	if err != nil {
		return nil, err
	}

	instance.db = db
	instance.logger = logger
	return instance, nil
}

func Disconnect() {
	if instance == nil || instance.db == nil {
		return
	}
	err := instance.db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	instance = nil
}

func (p *Postgres) GetConn(ctx context.Context) (*sqlx.Conn, error) {
	if instance == nil {
		return nil, fmt.Errorf("no Postgres instance found")
	}
	if ctx == nil {
		return nil, fmt.Errorf("no context")
	}
	conn, err := instance.getConnection(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (p *Postgres) getConnection(ctx context.Context) (*sqlx.Conn, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("context is already cancelled")
	default:
		conn, err := instance.db.Connx(ctx)
		if err != nil {
			return nil, fmt.Errorf("cannot get connection from pool: %v", err)
		}
		return conn, nil
	}
}

func (p *Postgres) ReturnConn(ctx context.Context, conn *sqlx.Conn) {
	select {
	case <-ctx.Done():
		return
	default:
		err := conn.Close()
		if err != nil {
			p.logger.Error(err.Error())
		}
	}
}
