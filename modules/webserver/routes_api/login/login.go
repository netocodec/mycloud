package login

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../../../db"
	"../../auth"
)

type LoginCredentials struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

func DoLogin(c *gin.Context) {
	var loginReq LoginCredentials

	jsonErr := c.BindJSON(&loginReq)
	if isValidLogin, loginUser := db.LoginMembership(loginReq.Username, loginReq.Password); isValidLogin && jsonErr == nil {
		tokenStr := auth.GenerateToken(db.UsersList{
			UserName:     loginUser.UserName,
			UserPassword: loginUser.UserPassword,
			IsAdmin:      loginUser.IsAdmin,
		})

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("mc_tok", tokenStr, 6600, "/", "25.38.61.125", false, true)
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
