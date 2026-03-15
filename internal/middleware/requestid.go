package middleware

import (
	"github.com/gin-gonic/gin"
)

// RequestIDMiddleware generates and sets a request ID in the context
// It also sets the X-Request-ID header in the response
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if request ID exists in header (from upstream services)
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			// Generate new request ID if not present
			requestID = generateRequestID()
		}

		// Set request ID in context
		ctx := setRequestIDToContext(c.Request.Context(), requestID)
		c.Request = c.Request.WithContext(ctx)

		// Set request ID in response header
		c.Header("X-Request-ID", requestID)

		// Continue processing
		c.Next()
	}
}
