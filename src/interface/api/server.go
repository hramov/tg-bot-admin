package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hramov/tg-bot-admin/src/config"
	"github.com/hramov/tg-bot-admin/src/interface/api/routes"
)

func StartServer() *gin.Engine {
	r := gin.New()

	// set release mode
	gin.SetMode(gin.ReleaseMode)

	// use cors middleware, access to everybody
	r.Use(cors.New(config.CorsConfig))

	// use recovery from panic middleware
	r.Use(gin.Recovery())
	// init routes
	api := r.Group("/api")

	auth := api.Group("/auth")
	routes.InitAuth(auth)

	user := api.Group("/user")
	routes.InitUser(user)

	return r
}
