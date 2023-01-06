package data_source

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres"
	"github.com/jmoiron/sqlx"
)

const (
	Postgres = iota
)

type DataSource interface {
	GetConn(ctx context.Context) (*sqlx.Conn, error)
}

func DatabaseFactory(dbName uint8) (DataSource, error) {
	switch dbName {
	case Postgres:
		return postgres.Connect()
	default:
		return nil, fmt.Errorf("no data source found")
	}
}
