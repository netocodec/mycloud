package webserver

import (
	"fmt"

	"../config"
	"./middleware"
	dashboardPage "./route/dashboard"
	indexPage "./route/index"
	myFilesPage "./route/myfiles"
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
	rootRouter.Use(middleware.NoCache())
	rootRouter.Static("/assets", "./static")
	rootRouter.LoadHTMLGlob("templates/**/*")

	// Login Page
	rootRouter.GET("/", indexPage.Index)

	memberRouter := rootRouter.Group("/member")
	memberRouter.Use(middleware.AuthorizePage())
	{
		memberRouter.GET("/dashboard", dashboardPage.Dashboard)
		memberRouter.GET("/myfiles", myFilesPage.MyFiles)
	}

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
			filesRouter.GET("/get/files", fshared.GetFolderContent)
			filesRouter.POST("/mk/:folder_name", fshared.MakeDir)
			filesRouter.POST("/upload/:file_name/:chunk_flag", fshared.UploadFiles)
		}
	}

	rootRouter.GET("/ping", ping.Ping)

	return rootRouter
}

func InitWebServer() {
	rootRouter := LoadWebServer()
	rootRouter.Run(fmt.Sprintf(":%d", config.GetServerPort()))
}
