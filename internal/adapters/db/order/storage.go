package order

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/domain/order"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type orderStorage struct {
	db     db.Connector
	logger *logging.Logger
}

func NewStorage(logger *logging.Logger, db db.Connector) order.IStorage {
	return &orderStorage{db: db, logger: logger}
}
