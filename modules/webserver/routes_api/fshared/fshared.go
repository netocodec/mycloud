package fshared

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../../../mfs"
	"../../auth"
)

type DirInfo struct {
	CurrentDir string `json:"currentd" binding:"required"`
}

func GetFolderContent(c *gin.Context) {
	folderName := c.Query("fname")
	userInfo := auth.GetHTTPToken(c)

	content := mfs.ContentInformation{
		ContentFullRoot: folderName,
		Type:            mfs.Directory,
	}

	if isSuccess, directoryInfo := mfs.GetContentOnUserCloud(userInfo.UserID, content); isSuccess {
		c.JSON(http.StatusOK, gin.H{
			"folder_content": directoryInfo.ContentData,
			"total_files":    directoryInfo.ContentSize,
			"msg_type":       "GFOLDER_SUCCESS",
		})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message":  "Cannot get this folder, try later!",
			"msg_type": "GFOLDER_FAIL",
		})
	}
}

func MakeDir(c *gin.Context) {
	dirName := c.Param("folder_name")
	userInfo := auth.GetHTTPToken(c)
	var dirInfo DirInfo

	c.BindJSON(&dirInfo)

	content := mfs.ContentInformation{
		ContentFullRoot: dirInfo.CurrentDir,
		ContentName:     dirName,
		Type:            mfs.Directory,
	}

	if isSuccess := mfs.CreateContentOnUserCloud(userInfo.UserID, content); isSuccess {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Folder has been created!",
			"msg_type": "MK_SUCESS",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Cannot create this folder!",
			"msg_type": "MK_FAIL",
		})
	}
}
