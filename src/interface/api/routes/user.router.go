package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hramov/tg-bot-admin/src/interface/api/controllers"
	"github.com/hramov/tg-bot-admin/src/interface/api/guards"
)

func InitUser(r *gin.RouterGroup) {
	r.GET("/info", guards.JwtGuard(), controllers.UserInfo)
}
