package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware restricts access based on the user role
func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetBool("is_admin")
		if !role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied: insufficient permissions"})
			return
		}
		c.Next()
	}
}
