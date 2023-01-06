package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hramov/tg-bot-admin/src/interface/api/controllers"
	"github.com/hramov/tg-bot-admin/src/interface/api/guards"
)

func InitAuth(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
	r.POST("/refresh", guards.JwtGuard(), controllers.Refresh)
}
