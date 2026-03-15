package logger

import (
	"go-api-project/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware creates and attaches a logger interface to the context
// It uses the request ID from context if available
func LoggerMiddleware(baseLogger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request ID from context (set by middleware.RequestIDMiddleware)
		ctx := c.Request.Context()
		requestID := middleware.GetRequestIDFromContext(ctx)

		// Create logger with request ID and request context
		logger := baseLogger.
			WithStr("request_id", requestID).
			WithStr("method", c.Request.Method).
			WithStr("path", c.Request.URL.Path).
			WithStr("remote_addr", c.ClientIP()).
			WithStr("user_agent", c.Request.UserAgent())

		// Set logger in context
		ctx = SetLoggerToContext(ctx, logger)
		c.Request = c.Request.WithContext(ctx)

		// Log request start
		logger.Info("Request started")

		// Measure request duration
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		// Get logger from context (might have been modified by handlers)
		responseLogger := GetLoggerFromContext(c.Request.Context())

		// Log request completion
		responseLogger.With(
			Str("status", string(rune(c.Writer.Status()))),
			Str("duration", duration.String()),
			Int("status_code", c.Writer.Status()),
		).Info("Request completed")

		// If there are errors, log them
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				responseLogger.WithErr(err.Err).Error("Request error")
			}
		}
	}
}
