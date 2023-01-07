package composite

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/api/user"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	user2 "github.com/hramov/tg-bot-admin/internal/adapters/db/user"
	user3 "github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/julienschmidt/httprouter"
)

func NewUser(pg db.Connector, router *httprouter.Router) {
	storage := user2.NewStorage(pg)
	service := user3.NewService(storage)
	handler := user.NewHandler(service)
	handler.Init(router)
}
