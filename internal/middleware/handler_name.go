package middleware

import (
	"github.com/gin-gonic/gin"
)

func HandlerName() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullName := c.FullPath()
		c.Set("handler_name", fullName)
		// Before function
		c.Next()
	}
}
