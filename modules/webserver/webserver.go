package webserver

import (
	"../mem"
	"./middleware"
	"./route/ping"
	ping_api "./routes_api/ping"

	"github.com/gin-gonic/gin"
)

func InitWebServer() {
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

		membershipRouter := apiRouter.Group("/membership")
		membershipRouter.Use(middleware.AuthorizeJWT())
		{
		}
	}

	rootRouter.GET("/ping", ping.Ping)

	rootRouter.Run(":8080")
}
