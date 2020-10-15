package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "main_page.tmpl", gin.H{
		"title": "Homepage",
		"page":  "index",
	})
}
