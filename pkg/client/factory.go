package client

import (
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/pkg/client/postgres"
)

const (
	Postgres = iota
)

func DatabaseFactory(dbName uint8) (db.Connector, error) {
	switch dbName {
	case Postgres:
		return postgres.Connect()
	default:
		return nil, fmt.Errorf("no data source found")
	}
}
