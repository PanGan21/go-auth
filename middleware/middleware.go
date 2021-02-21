package middleware

import (
	"go-auth/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware returns the authentication middleware
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
