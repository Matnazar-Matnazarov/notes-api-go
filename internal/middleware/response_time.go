package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ResponseTime middleware adds response time to all API responses
// This helps track how long each request takes to process
func ResponseTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record start time and store it in context
		// This allows handlers to calculate response time before sending response
		start := time.Now()
		c.Set("start_time", start)

		// Process request
		c.Next()

		// Calculate response time after handler completes
		duration := time.Since(start)

		// Add response time to response header
		c.Header("X-Response-Time", duration.String())
	}
}
