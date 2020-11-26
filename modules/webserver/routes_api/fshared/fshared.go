package fshared

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"../../../mfs"
	"../../auth"
)

type DirInfo struct {
	CurrentDir string `json:"currentd" binding:"required"`
}

type UploadData struct {
	CurrentDir string `json:"currentd" binding:"required"`
	Chunk      string `json:"chunk" binding:"required"`
}

var uploadPool map[string]*os.File = make(map[string]*os.File)

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
			"msg_type": "MK_SUCCESS",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Cannot create this folder!",
			"msg_type": "MK_FAIL",
		})
	}
}

func UploadFiles(c *gin.Context) {
	filename := c.Param("file_name")
	chunkFlag := c.Param("chunk_flag")
	userInfo := auth.GetHTTPToken(c)

	var uploadData UploadData

	c.BindJSON(&uploadData)

	fileOp, fileExists := uploadPool[filename]
	if !fileExists && chunkFlag == "1" {
		isSuccess := false

		if fileOp, fileOpErr := mfs.OpenFileStream(uploadData.CurrentDir, filename, userInfo.UserID); fileOpErr == nil {
			if bdata, berr := base64.StdEncoding.DecodeString(uploadData.Chunk); berr == nil {
				isSuccess = true
				fileOp.Write(bdata)
			}
			fileOp.Close()
		}

		if isSuccess {
			c.JSON(http.StatusOK, gin.H{
				"message":  fmt.Sprintf("File %s uploaded with success!", filename),
				"msg_type": "UPLOAD_SUCCESS",
			})
		} else {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message":  fmt.Sprintf("Cannot upload file %s!", filename),
				"msg_type": "UPLOAD_FAIL",
			})
		}
	} else if !fileExists && chunkFlag == "0" {
		fileOp, fileOpErr := mfs.OpenFileStream(uploadData.CurrentDir, filename, userInfo.UserID)

		if fileOpErr == nil {
			uploadPool[filename] = fileOp

			if bdata, berr := base64.StdEncoding.DecodeString(uploadData.Chunk); berr == nil {
				fileOp.Write(bdata)
			}

			c.JSON(http.StatusOK, gin.H{
				"message":  fmt.Sprintf("File %s uploaded with success!", filename),
				"msg_type": "UPLOAD_SUCCESS",
			})
		} else {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message":  fmt.Sprintf("Cannot upload file %s! (Cannot open file)", filename),
				"msg_type": "UPLOAD_OPEN_FAIL",
			})
		}
	} else if fileExists {
		if bdata, berr := base64.StdEncoding.DecodeString(uploadData.Chunk); berr == nil {
			fileOp.Write(bdata)
		}

		if chunkFlag == "1" {
			fileOp.Close()
			delete(uploadPool, filename)

			c.JSON(http.StatusOK, gin.H{
				"message":  fmt.Sprintf("File %s uploaded with success!", filename),
				"msg_type": "UPLOAD_SUCCESS",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":  "Chunk uploaded!",
				"msg_type": "UPLOAD_CHUNK_SUCCESS",
			})
		}
	}
}
