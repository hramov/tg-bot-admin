package db

import (
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/pkg/db/postgres"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

const (
	Postgres = iota
)

func DatabaseFactory(dbName uint8, cfg *config.Config, logger *logging.Logger) (db.Connector, error) {
	switch dbName {
	case Postgres:
		return postgres.Connect(cfg, logger)
	default:
		return nil, fmt.Errorf("no data source found")
	}
}
