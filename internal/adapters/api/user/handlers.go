package user

import (
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/api/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := utils.GetBody[user.LoginDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}
	serviceResponse, serviceError := h.service.Login(r.Context(), &body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, serviceResponse, w)
}

func (h *handler) Refresh(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := utils.GetBody[user.LoginResponseDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}
	serviceResponse, serviceError := h.service.Refresh(r.Context(), &body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, serviceResponse, w)
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := utils.GetBody[user.CreateDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}
	serviceResponse, serviceError := h.service.Create(r.Context(), &body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, serviceResponse, w)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users, err := h.service.GetAll(r.Context(), 10, 0)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, users, w)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	rawId := params.ByName("user_id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "wrong id format", w)
		return
	}
	users, err := h.service.Delete(r.Context(), id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, users, w)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := utils.GetBody[user.UpdateDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}
	serviceResponse, serviceError := h.service.Update(r.Context(), &body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, serviceResponse, w)
}

func (h *handler) GetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	rawId := params.ByName("user_id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "wrong id format", w)
		return
	}
	users, err := h.service.GetById(r.Context(), id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, users, w)
}
