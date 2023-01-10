package db

import (
	"fmt"
	"github.com/hramov/tg-bot-admin/pkg/db/postgres"
)

const (
	Postgres = iota
)

func DatabaseFactory(dbName uint8, cfg postgres.Config) (Connector, error) {
	switch dbName {
	case Postgres:
		return postgres.Connect(cfg)
	default:
		return nil, fmt.Errorf("no data source found")
	}
}
