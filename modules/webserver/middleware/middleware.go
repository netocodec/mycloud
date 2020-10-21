package middleware

import (
	"net/http"

	"../auth"

	"github.com/gin-gonic/gin"
)

const authHeader string = "Bearer "

func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, max-age=0")
		c.Next()
	}
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieData, cookieErr := c.Cookie("mc_tok")

		if cookieErr == nil {
			if hasError, _ := auth.DecodeToken(cookieData); hasError {
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
