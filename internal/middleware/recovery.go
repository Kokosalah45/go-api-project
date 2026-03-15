package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware recovers from panics and returns a proper JSON response
func RecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		var err error
		if e, ok := recovered.(error); ok {
			err = e
		} else if s, ok := recovered.(string); ok {
			err = &panicError{msg: s}
		} else {
			err = &panicError{msg: "Unknown panic"}
		}

		// Return JSON response with error details
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":       "Internal Server Error",
			"request_id":  GetRequestIDFromContext(c.Request.Context()),
			"panic_error": err.Error(),
		})
	})
}

// panicError is a simple error type for panics
type panicError struct {
	msg string
}

func (e *panicError) Error() string {
	return e.msg
}
