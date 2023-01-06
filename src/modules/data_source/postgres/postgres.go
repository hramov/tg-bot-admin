package postgres

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/src/config"
	dbInit "github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init"
	"github.com/hramov/tg-bot-admin/src/modules/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Postgres struct {
	db *sqlx.DB
}

var instance *Postgres

func Connect() (*Postgres, error) {
	if instance != nil {
		return instance, nil
	}

	instance = &Postgres{}
	db, err := sqlx.Open("postgres", os.Getenv("PG_DSN"))
	if err != nil {
		logger.Instance.Error(err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	err = dbInit.Start(db)
	if err != nil {
		logger.Instance.Error(err.Error())
		return nil, err
	}

	instance.db = db
	return instance, nil
}

func Disconnect() {
	if instance == nil || instance.db == nil {
		return
	}
	err := instance.db.Close()
	if err != nil {
		logger.Instance.Error(err.Error())
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
	conn, err := instance.db.Connx(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get connection from pool: %v", err)
	}
	return conn, nil
}
