package webserver

import (
	"fmt"

	"../config"
	"./middleware"
	indexPage "./route/index"
	"./route/ping"
	"./routes_api/fshared"
	"./routes_api/login"
	ping_api "./routes_api/ping"
	"./routes_api/settings"

	"github.com/gin-gonic/gin"
)

func LoadWebServer() *gin.Engine {
	mode := gin.ReleaseMode

	if config.DebugMode {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	rootRouter := gin.New()
	rootRouter.Use(gin.Recovery())
	rootRouter.Static("/assets", "./static")
	rootRouter.LoadHTMLGlob("templates/**/*")

	rootRouter.GET("/", indexPage.Index)

	apiRouter := rootRouter.Group("/api")
	{
		apiRouter.GET("/ping", ping_api.Ping)
		apiRouter.POST("/login", login.DoLogin)

		membershipRouter := apiRouter.Group("/membership")
		membershipRouter.Use(middleware.AuthorizeJWT())
		{
			membershipRouter.GET("/settings", settings.GetUserSettings)
			membershipRouter.POST("/settings", settings.EditUserSettings)
		}

		filesRouter := apiRouter.Group("/fshared")
		filesRouter.Use(middleware.AuthorizeJWT())
		{
			filesRouter.POST("/get/:folder_name", fshared.GetFolderContent)
			filesRouter.PUT("/mk/:folder_name", fshared.MakeDir)
		}
	}

	rootRouter.GET("/ping", ping.Ping)

	return rootRouter
}

func InitWebServer() {
	rootRouter := LoadWebServer()
	rootRouter.Run(fmt.Sprintf(":%d", config.GetServerPort()))
}
