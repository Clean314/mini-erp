package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "admin-token" {
			c.Set("userID", uint(1))
			c.Set("role", "ADMIN")
		} else if token == "manager-token" {
			c.Set("userID", uint(2))
			c.Set("role", "MANAGER")
		} else {
			c.Set("userID", uint(3))
			c.Set("role", "USER")
		}

		c.Next()
	}
}