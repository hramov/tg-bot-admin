package user

import (
	"github.com/go-chi/chi/v5"
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
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
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

// Refresh
// @Summary Refresh token pair handler
// @Tags Refresh
// @Accept       json
// @Produce      json
// @Param loginResponse body user.LoginResponseDto true "access and refresh tokens"
// @Success      200  {object}  user.LoginResponseDto
// @Failure 401
// @Router /api/refresh [post]
func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
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

// Register
// @Summary Register user
// @Tags Register
// @Accept       json
// @Produce      json
// @Param createUserDto body user.CreateDto true "user data for register"
// @Success      200
// @Failure 400
// @Router /api/register [post]
func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
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

// Get
// @Summary Get users
// @Tags Get
// @Accept       json
// @Produce      json
// @Param        count   query      int  false  "Limit"
// @Param        start   query      int  false  "Offset"
// @Param        sortBy   query      string  false  "Order By"
// @Param        desc   query      boolean  false  "true = desc"
// @Success      200  {array}  user.User
// @Failure 401
// @Failure 500
// @Router /api/users [get]
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll(r.Context())
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
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
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
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
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
func (h *handler) GetOne(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "user_id")
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
