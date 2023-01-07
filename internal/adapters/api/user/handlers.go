package user

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/api/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
	service user.IService
}

const (
	loginUrl    = "/api/user/login"
	registerUrl = "/api/user/register"
	usersUrl    = "/api/user"
)

func NewHandler(service user.IService) api.Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Init(router *httprouter.Router) {
	router.GET(loginUrl, h.Login)
	router.GET(registerUrl, h.Register)
	router.GET(usersUrl, h.Get)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {}

func (h *handler) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {}

func (h *handler) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users, err := h.service.GetAll(r.Context(), 10, 0)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, users, w)
}
