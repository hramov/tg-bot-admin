package composite

import (
	"github.com/go-playground/validator/v10"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/handlers/order"
	product2 "github.com/hramov/tg-bot-admin/internal/adapters/api/handlers/product"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/handlers/user"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	order3 "github.com/hramov/tg-bot-admin/internal/adapters/db/order"
	product3 "github.com/hramov/tg-bot-admin/internal/adapters/db/product"
	user2 "github.com/hramov/tg-bot-admin/internal/adapters/db/user"
	"github.com/hramov/tg-bot-admin/internal/config"
	order2 "github.com/hramov/tg-bot-admin/internal/domain/order"
	"github.com/hramov/tg-bot-admin/internal/domain/product"
	user3 "github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type UserComposite struct{}

func (uc *UserComposite) Register(logger *logging.Logger, cfg *config.Config, pg db.Connector, router *httprouter.Router) {
	storage := user2.NewStorage(logger, pg)
	service := user3.NewService(storage, validator.New(), logger, cfg)
	handler := user.NewHandler(logger, service)
	handler.Init(router)
}

type ProductComposite struct{}

func (uc *ProductComposite) Register(logger *logging.Logger, cfg *config.Config, pg db.Connector, router *httprouter.Router) {
	storage := product3.NewStorage(logger, pg)
	service := product.NewService(storage, validator.New(), logger, cfg)
	handler := product2.NewHandler(logger, service)
	handler.Init(router)
}

type OrderComposite struct{}

func (uc *OrderComposite) Register(logger *logging.Logger, cfg *config.Config, pg db.Connector, router *httprouter.Router) {
	storage := order3.NewStorage(logger, pg)
	service := order2.NewService(storage, validator.New(), logger, cfg)
	handler := order.NewHandler(logger, service)
	handler.Init(router)
}
