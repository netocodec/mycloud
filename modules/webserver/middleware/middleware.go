package middleware

import (
	"net/http"

	"../auth"

	"github.com/gin-gonic/gin"
)

const authHeader string = "Bearer "

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization != "" {
			tokenStr := authorization[len(authHeader):]

			if hasError, _ := auth.DecodeToken(tokenStr); hasError {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func AuthorizePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieData, cookieErr := c.Cookie("mc_tok")

		if cookieErr == nil {
			if hasError, _ := auth.DecodeToken(cookieData); hasError {
				c.Redirect(http.StatusPermanentRedirect, "/?err=ACCOUNT_NOT_AUTH")
				c.Abort()
			}
		} else {
			c.Redirect(http.StatusPermanentRedirect, "/?err=ACCOUNT_NOT_AUTH")
			c.Abort()
		}
	}
}
