package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CaptureIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		c.Set("ip_address", ip)
		c.Next()
	}
}
