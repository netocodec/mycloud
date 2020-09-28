package login

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../../../db"
	"../../auth"
)

func DoLogin(c *gin.Context) {
	username := c.PostForm("user")
	password := c.PostForm("pass")

	if isValidLogin, loginUser := db.LoginMembership(username, password); isValidLogin {
		tokenStr := auth.GenerateToken(db.UsersList{
			UserName:     loginUser.UserName,
			UserPassword: loginUser.UserPassword,
			IsAdmin:      loginUser.IsAdmin,
		})

		c.JSON(http.StatusOK, gin.H{
			"token": tokenStr,
		})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message":  "Username and password not valid!",
			"msg_type": "LOGIN_FAIL",
		})
	}

}
