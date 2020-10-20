package dashboard

import (
	"net/http"

	"../../menu"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "main_page.tmpl", gin.H{
		"title":      "Dashboard",
		"page":       "dashboard",
		"isAuthPage": true,
		"menu":       menu.GetMenu(),
	})
}
