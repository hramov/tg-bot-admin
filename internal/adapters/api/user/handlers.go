package user

import (
	"github.com/hramov/tg-bot-admin/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
	service Service
}

const (
	loginUrl    = "/api/user/login"
	registerUrl = "/api/user/register"
)

func NewHandler(service Service) api.Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Init(router *httprouter.Router) {
	router.GET(loginUrl, h.Login)
	router.GET(registerUrl, h.Register)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//body, err := utils.GetBody[auth.Login](c)
	//if err != nil {
	//	utils.SendError(http.StatusBadRequest, "Cannot parse request body", c)
	//	return
	//}
	//
	//login := auth.Login{
	//	Email:    body.Email,
	//	Password: body.Password,
	//}
	//at, rt, lErr := login.Login(ctx)
	//if lErr != nil {
	//	utils.SendError(lErr.Status(), lErr.Error(), c)
	//	return
	//}
}

//ctx, cancel := customContext.WithCancel()
//go utils.MaintainRequest(c.Request.Context(), cancel)
//
//body, err := utils.GetBody[auth.Refresh](c)
//if err != nil {
//utils.SendError(http.StatusBadRequest, "Cannot parse request body", c)
//return
//}
//
//refresh := auth.Refresh{
//AccessToken:  body.AccessToken,
//RefreshToken: body.RefreshToken,
//}
//
//at, rt, rErr := refresh.Refresh(ctx)
//if err != nil {
//utils.SendError(rErr.Status(), rErr.Error(), c)
//return
//}
//
//utils.SendResponse(http.StatusOK, &gin.H{
//"access_token":  at,
//"refresh_token": rt,
//}, c)

func (h *handler) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
