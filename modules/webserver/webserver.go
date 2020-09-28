package webserver

import (
	"../mem"
	"./middleware"
	"./route/ping"
	"./routes_api/login"
	ping_api "./routes_api/ping"

	"github.com/gin-gonic/gin"
)

func LoadWebServer() *gin.Engine {
	mode := gin.ReleaseMode

	if mem.DebugMode {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	rootRouter := gin.New()

	rootRouter.Use(gin.Recovery())

	apiRouter := rootRouter.Group("/api")
	{
		apiRouter.GET("/ping", ping_api.Ping)
		apiRouter.POST("/login", login.DoLogin)

		membershipRouter := apiRouter.Group("/membership")
		membershipRouter.Use(middleware.AuthorizeJWT())
		{
		}
	}

	rootRouter.GET("/ping", ping.Ping)

	return rootRouter
}

func InitWebServer() {
	rootRouter := LoadWebServer()
	rootRouter.Run(":8080")
}
