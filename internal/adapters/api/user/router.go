package user

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/api/guards"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	service user.IService
}

const (
	loginUrl    = "/api/login"
	registerUrl = "/api/register"
	usersUrl    = "/api/users"
	userUrl     = "/api/user/:user_id"
)

func NewHandler(service user.IService) api.Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Init(router *httprouter.Router) {
	router.POST(loginUrl, h.Login)
	router.GET(registerUrl, h.Register)
	router.GET(usersUrl, guards.JwtGuard(h.Get, []string{"admin"}))
	router.GET(userUrl, guards.JwtGuard(h.GetOne, []string{"admin", "equal_id"}))
	router.PUT(userUrl, guards.JwtGuard(h.Update, []string{"admin", "equal_id"}))
	router.DELETE(userUrl, guards.JwtGuard(h.Delete, []string{"admin", "equal_id"}))
}
