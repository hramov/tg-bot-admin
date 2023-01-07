package composite

import (
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/user"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	user2 "github.com/hramov/tg-bot-admin/internal/adapters/db/user"
	user3 "github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func NewUser(pg db.Connector, router *httprouter.Router, logger *logging.Logger) {
	storage := user2.NewStorage(pg)
	service := user3.NewService(storage, validator.New(), logger)
	handler := user.NewHandler(service)
	handler.Init(router)
}
