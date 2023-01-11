package user

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/middlewares"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
	service user.Service
	logger  *logging.Logger
}

const (
	loginUrl    = "/api/login"
	registerUrl = "/api/register"
	refreshUrl  = "/api/refresh"
	usersUrl    = "/api/users"
	userUrl     = "/api/user/:user_id"
)

func NewHandler(logger *logging.Logger, service user.Service) api.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Init(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, registerUrl, h.Register)
	router.HandlerFunc(http.MethodGet, usersUrl, middlewares.Auth(middlewares.Filter(h.Get), []string{"admin"}))
	router.HandlerFunc(http.MethodGet, userUrl, middlewares.Auth(h.GetOne, []string{"admin", "equal_id"}))
	router.HandlerFunc(http.MethodPost, loginUrl, h.Login)
	router.HandlerFunc(http.MethodPost, refreshUrl, h.Refresh)
	router.HandlerFunc(http.MethodPut, userUrl, middlewares.Auth(h.Update, []string{"admin", "equal_id"}))
	router.HandlerFunc(http.MethodDelete, userUrl, middlewares.Auth(h.Delete, []string{"admin", "equal_id"}))
}
