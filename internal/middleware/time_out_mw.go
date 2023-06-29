package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Update the request context to the one with timeout
		c.Request = c.Request.WithContext(ctx)

		c.Next()

	}
}
