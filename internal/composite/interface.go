package composite

import (
	"github.com/go-chi/chi/v5"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type Composite interface {
	Register(logger *logging.Logger, cfg *config.Config, pg db.Connector, router *chi.Mux)
}
