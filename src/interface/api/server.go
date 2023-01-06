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
	// serve frontend
	//_, b, _, _ := runtime.Caller(0)
	//basePath := filepath.Dir(b)
	//
	//folder := basePath + "/public"
	//file := basePath + "/public/index.html"
	//
	//_ = mime.AddExtensionType(".js", "text/javascript")
	//_ = mime.AddExtensionType(".css", "text/css")
	//r.StaticFS("/public", gin.Dir(folder, true))
	//
	//r.LoadHTMLFiles(file)
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/dashboard")
	//})
	//
	//r.GET("/dashboard/*page", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})

	// init routes
	api := r.Group("/api")

	auth := api.Group("/auth")
	routes.InitAuth(auth)

	return r
}
