package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/middlewares"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/logging"
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
	userUrl     = "/api/user/{user_id}"
)

func NewHandler(logger *logging.Logger, service user.Service) api.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Init(router *chi.Mux) {
	router.MethodFunc(http.MethodPost, registerUrl, middlewares.Timeout(h.Register))
	router.MethodFunc(http.MethodGet, usersUrl, middlewares.Timeout(middlewares.Auth(middlewares.Filter(h.Get), []string{"admin"})))
	router.MethodFunc(http.MethodGet, userUrl, middlewares.Timeout(middlewares.Auth(h.GetOne, []string{"admin", "equal_id"})))
	router.MethodFunc(http.MethodPost, loginUrl, middlewares.Timeout(h.Login))
	router.MethodFunc(http.MethodPost, refreshUrl, middlewares.Timeout(h.Refresh))
	router.MethodFunc(http.MethodPut, userUrl, middlewares.Timeout(middlewares.Auth(h.Update, []string{"admin", "equal_id"})))
	router.MethodFunc(http.MethodDelete, userUrl, middlewares.Timeout(middlewares.Auth(h.Delete, []string{"admin", "equal_id"})))
}
