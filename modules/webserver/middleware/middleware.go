package middleware

import (
	"net/http"

	"../auth"

	"github.com/gin-gonic/gin"
)

const authHeader string = "Bearer"

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		tokenStr := authorization[len(authHeader):]

		if hasError, _ := auth.DecodeToken(tokenStr); hasError {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
