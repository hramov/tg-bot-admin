package user

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/http/guards"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	service user.IService
	logger  *logging.Logger
}

const (
	loginUrl    = "/api/login"
	registerUrl = "/api/register"
	refreshUrl  = "/api/refresh"
	usersUrl    = "/api/users"
	userUrl     = "/api/user/:user_id"
)

func NewHandler(logger *logging.Logger, service user.IService) api.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Init(router *httprouter.Router) {
	router.GET(registerUrl, h.Register)
	router.GET(usersUrl, guards.JwtGuard(h.Get, []string{"admin"}))
	router.GET(userUrl, guards.JwtGuard(h.GetOne, []string{"admin", "equal_id"}))
	router.POST(loginUrl, h.Login)
	router.POST(refreshUrl, h.Refresh)
	router.PUT(userUrl, guards.JwtGuard(h.Update, []string{"admin", "equal_id"}))
	router.DELETE(userUrl, guards.JwtGuard(h.Delete, []string{"admin", "equal_id"}))
}
