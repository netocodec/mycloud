package myfiles

import (
	"net/http"

	"../../menu"
	"github.com/gin-gonic/gin"
)

func MyFiles(c *gin.Context) {
	c.HTML(http.StatusOK, "main_page.tmpl", gin.H{
		"title":      "My Files",
		"page":       "myfiles",
		"isAuthPage": true,
		"menu":       menu.GetMenu(),
	})
}
