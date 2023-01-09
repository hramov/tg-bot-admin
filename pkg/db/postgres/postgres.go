package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Postgres struct {
	db *sqlx.DB
}

var instance *Postgres

func Connect(cfg Config) (*Postgres, error) {
	if instance != nil {
		return instance, nil
	}

	instance = &Postgres{}
	db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SslMode))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenCons)
	db.SetMaxIdleConns(cfg.MaxIdleCons)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err != nil {
		return nil, err
	}

	instance.db = db
	return instance, nil
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

func (p *Postgres) ReturnConn(ctx context.Context, conn *sqlx.Conn) {
	select {
	case <-ctx.Done():
		return
	default:
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}
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
