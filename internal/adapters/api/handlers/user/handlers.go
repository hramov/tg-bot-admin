package user

import (
	"context"
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Login
// @Summary Login handler
// @Tags Login
// @Accept       json
// @Produce      json
// @Param login body user.LoginDto true "login credentials"
// @Success      200  {object}  user.LoginResponseDto
// @Failure 401
// @Router /api/login [post]
func (h *handler) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := utils.GetBody[user.LoginDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), config.DefaultTimeout)
	defer cancel()

	serviceResponse, serviceError := h.service.Login(ctx, &body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, serviceResponse, w)
}

// Refresh
// @Summary Refresh token pair handler
// @Tags Refresh
// @Accept       json
// @Produce      json
// @Param loginResponse body user.LoginResponseDto true "access and refresh tokens"
// @Success      200  {object}  user.LoginResponseDto
// @Failure 401
// @Router /api/refresh [post]
func (h *handler) Refresh(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := utils.GetBody[user.LoginResponseDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), config.DefaultTimeout)
	defer cancel()

	serviceResponse, serviceError := h.service.Refresh(ctx, &body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, serviceResponse, w)
}

// Register
// @Summary Register user
// @Tags Register
// @Accept       json
// @Produce      json
// @Param createUserDto body user.CreateDto true "user data for register"
// @Success      200
// @Failure 400
// @Router /api/register [post]
func (h *handler) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := utils.GetBody[user.CreateDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), config.DefaultTimeout)
	defer cancel()

	serviceResponse, serviceError := h.service.Create(ctx, &body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, serviceResponse, w)
}

// Get
// @Summary Get users
// @Tags Get
// @Accept       json
// @Produce      json
// @Param        limit   path      int  true  "Limit"
// @Param        offset   path      int  true  "Offset"
// @Success      200  {array}  user.User
// @Failure 401
// @Failure 500
// @Router /api/users/:limit/:offset [get]
func (h *handler) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), config.DefaultTimeout)
	defer cancel()

	users, err := h.service.GetAll(ctx)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, users, w)
}

// Delete
// @Summary Delete user
// @Tags Delete
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {array}  user.User
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/user/:id [delete]
func (h *handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	rawId := params.ByName("user_id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "wrong id format", w)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), config.DefaultTimeout)
	defer cancel()

	users, err := h.service.Delete(ctx, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, users, w)
}

// Update
// @Summary Update user
// @Tags Update
// @Accept       json
// @Produce      json
// @Param        updateUserDto body user.UpdateDto true "user data for update"
// @Success      201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/user/:id [put]
func (h *handler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := utils.GetBody[user.UpdateDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), config.DefaultTimeout)
	defer cancel()

	serviceResponse, serviceError := h.service.Update(ctx, &body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusCreated, serviceResponse, w)
}

// GetOne
// @Summary GetOne user
// @Tags GetOne
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  user.User
// @Failure 401
// @Failure 500
// @Router /api/user/:id [get]
func (h *handler) GetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	rawId := params.ByName("user_id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "wrong id format", w)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), config.DefaultTimeout)
	defer cancel()

	users, err := h.service.GetById(ctx, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, users, w)
}
