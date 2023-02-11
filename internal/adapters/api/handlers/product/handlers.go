package product

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hramov/tg-bot-admin/internal/domain/product"
	"github.com/hramov/tg-bot-admin/pkg/utils"
)

// Get
// @Summary Get products
// @Tags Get
// @Accept       json
// @Produce      json
// @Param        count   query      int  true  "Limit"
// @Param        start   query      int  true  "Offset"
// @Param        sortBy   query      string  false  "Order By"
// @Param        desc   query      boolean  false  "true = desc"
// @Success      200  {array}  product.Product
// @Failure 401
// @Failure 500
// @Router /api/products [get]
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		utils.SendError(http.StatusBadRequest, "param count has wrong format", w)
		return
	}

	lastId, err := strconv.Atoi(r.URL.Query().Get("last_id"))
	if err != nil {
		utils.SendError(http.StatusBadRequest, "param last_id has wrong format", w)
		return
	}

	products, err := h.service.GetAll(r.Context(), limit, lastId)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, products, w)
}

// GetById
// @Summary Get product by id
// @Tags Get
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  product.Product
// @Failure 401
// @Failure 500
// @Router /api/products/:product_id [get]
func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "product_id")

	products, err := h.service.GetBy(r.Context(), "id", rawId)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, products, w)
}

// Create
// @Summary Create new product
// @Tags Create
// @Accept       json
// @Produce      json
// @Param        product	body	product.Product	true	"Product dto"
// @Success      200
// @Failure	401
// @Failure 500
// @Router /api/products [post]
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := utils.GetBody[product.InputWeedProduct](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}

	serviceResponse, serviceError := h.service.Create(r.Context(), body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusCreated, serviceResponse, w)
}

// Update
// @Summary Update product
// @Tags Update
// @Accept       json
// @Produce      json
// @Param        product   body	product.Product	true	"Product dto"
// @Success      200
// @Failure 401
// @Failure 500
// @Router /api/products [put]
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	body, err := utils.GetBody[product.InputWeedProduct](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "cannot parse body", w)
		return
	}

	serviceResponse, serviceError := h.service.Update(r.Context(), body)
	if serviceError != nil {
		utils.SendError(serviceError.Status(), serviceError.Error(), w)
		return
	}
	utils.SendResponse(http.StatusCreated, serviceResponse, w)
}

// Delete
// @Summary Delete product by id
// @Tags Delete
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200
// @Failure 401
// @Failure 500
// @Router /api/products/:product_id [delete]
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "product_id")

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "wrong product id", w)
		return
	}

	products, err := h.service.Delete(r.Context(), id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, products, w)
}
