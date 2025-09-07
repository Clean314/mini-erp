package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")
		for _, r := range requiredRoles {
			if userRole == r {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		c.Abort()
	}
}
