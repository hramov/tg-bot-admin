package product

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/hramov/tg-bot-admin/internal/domain/product"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	service product.IService
	logger  *logging.Logger
}

const ()

func NewHandler(logger *logging.Logger, service product.IService) api.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Init(router *httprouter.Router) {
}
