package product

import (
	"github.com/go-chi/chi/v5"
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/hramov/tg-bot-admin/internal/domain/product"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type handler struct {
	service product.Service
	logger  *logging.Logger
}

const ()

func NewHandler(logger *logging.Logger, service product.Service) api.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Init(router *chi.Mux) {
}
