package settings

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../../../db"
	"../../auth"
)

type EditSettings struct {
	UserPass             string `json:"pass"`
	UserPassConfirmation string `json:"passc"`
}

func GetUserSettings(c *gin.Context) {
	var userInfo db.UsersList = auth.GetHTTPToken(c)
	c.JSON(http.StatusOK, userInfo)
}

func EditUserSettings(c *gin.Context) {
	var userEdit EditSettings
	var userInfo db.UsersList = auth.GetHTTPToken(c)
	var statusCode int = http.StatusOK
	var statusResult gin.H = gin.H{
		"message":  "Information editted with success!",
		"msg_type": "INFO_EDIT_SUCESS",
	}

	if userJSONErr := c.BindJSON(&userEdit); userJSONErr != nil {
		statusCode = http.StatusNotAcceptable
		statusResult = gin.H{
			"message":  "Invalid fields!",
			"msg_type": "INVALID_FIELDS",
		}
	} else {
		if userEdit.UserPass != userEdit.UserPassConfirmation {
			statusCode = http.StatusNotAcceptable
			statusResult = gin.H{
				"message":  "Invalid confirmation fields!",
				"msg_type": "INVALID_CONFIRMATION_FIELDS",
			}
		} else {
			if hasChanged := db.EditUserPass(userEdit.UserPass, userInfo.UserID); !hasChanged {
				statusCode = http.StatusNotAcceptable
				statusResult = gin.H{
					"message":  "Cannot edit the password, try later!",
					"msg_type": "EDIT_FAILED",
				}
			}
		}
	}

	c.JSON(statusCode, statusResult)
}
