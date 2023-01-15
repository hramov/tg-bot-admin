package product

import (
	"github.com/hramov/tg-bot-admin/pkg/utils"
	"net/http"
	"strconv"
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
// @Success      200  {array}  user.User
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
