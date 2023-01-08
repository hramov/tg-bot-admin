package composite

import (
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/user"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	user2 "github.com/hramov/tg-bot-admin/internal/adapters/db/user"
	"github.com/hramov/tg-bot-admin/internal/config"
	user3 "github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type UserComposite struct{}

func (uc *UserComposite) Register(logger *logging.Logger, cfg *config.Config, pg db.Connector, router *httprouter.Router) {
	storage := user2.NewStorage(pg)
	service := user3.NewService(storage, validator.New(), logger, cfg)
	handler := user.NewHandler(service)
	handler.Init(router)
}
