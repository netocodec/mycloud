package webserver

import (
	"../mem"
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
	}

	rootRouter.GET("/ping", ping.Ping)

	rootRouter.Run(":8080")
}
