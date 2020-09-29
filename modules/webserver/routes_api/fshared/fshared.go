package fshared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IN_DEVELOPMENT FOR THE FUTURE
func GetFolderContent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"folder_content": "[]",
		"total_files":    0,
	})
}

// IN_DEVELOPMENT FOR THE FUTURE
func MakeDir(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":  "Folder has been created!",
		"msg_type": "MK_SUCESS",
	})
}
