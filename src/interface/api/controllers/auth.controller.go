package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hramov/tg-bot-admin/src/domain/auth"
	"github.com/hramov/tg-bot-admin/src/interface/api/utils"
	customContext "github.com/hramov/tg-bot-admin/src/modules/context"
	"net/http"
)

func Login(c *gin.Context) {
	ctx, cancel := customContext.WithCancel()
	go utils.MaintainRequest(c.Request.Context(), cancel)
	body, err := utils.GetBody[auth.Login](c)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "Cannot parse request body", c)
		return
	}

	login := auth.Login{
		Email:    body.Email,
		Password: body.Password,
	}
	at, rt, lErr := login.Login(ctx)
	if lErr != nil {
		utils.SendError(lErr.Status(), lErr.Error(), c)
		return
	}

	utils.SendResponse(http.StatusOK, &gin.H{
		"access_token":  at,
		"refresh_token": rt,
	}, c)
}

func Refresh(c *gin.Context) {
	ctx, cancel := customContext.WithCancel()
	go utils.MaintainRequest(c.Request.Context(), cancel)

	body, err := utils.GetBody[auth.Refresh](c)
	if err != nil {
		utils.SendError(http.StatusBadRequest, "Cannot parse request body", c)
		return
	}

	refresh := auth.Refresh{
		AccessToken:  body.AccessToken,
		RefreshToken: body.RefreshToken,
	}

	at, rt, rErr := refresh.Refresh(ctx)
	if err != nil {
		utils.SendError(rErr.Status(), rErr.Error(), c)
		return
	}

	utils.SendResponse(http.StatusOK, &gin.H{
		"access_token":  at,
		"refresh_token": rt,
	}, c)
}
