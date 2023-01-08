package composite

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type Composite interface {
	Register(logger *logging.Logger, cfg *config.Config, pg db.Connector, router *httprouter.Router)
}
