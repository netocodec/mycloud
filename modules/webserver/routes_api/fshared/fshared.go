package fshared

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IN_DEVELOPMENT FOR THE FUTURE
func GetFolderContent(c *gin.Context) {
	folderName := c.Query("fname")
	fmt.Println("FOLDER NAME: ", folderName)
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
