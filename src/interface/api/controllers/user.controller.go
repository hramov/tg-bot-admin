package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hramov/tg-bot-admin/src/domain/user"
	"github.com/hramov/tg-bot-admin/src/interface/api/utils"
	customContext "github.com/hramov/tg-bot-admin/src/modules/context"
	"net/http"
)

func UserInfo(c *gin.Context) {
	ctx, cancel := customContext.WithCancel()
	go utils.MaintainRequest(c.Request.Context(), cancel)

	id, ok := c.Get("id")
	if !ok {
		utils.SendError(http.StatusUnauthorized, "unauthorized", c)
		return
	}

	userData := user.User{
		Id: id.(int),
	}
	data, err := userData.Info(ctx)
	if err != nil {
		utils.SendError(err.Status(), err.Error(), c)
		return
	}

	utils.SendResponse(http.StatusOK, data, c)
}
