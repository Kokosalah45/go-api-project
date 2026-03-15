package middleware

import (
	"context"
	"github.com/google/uuid"
)

// requestIDKey is the key used to store the request ID in the context
type requestIDKey string

// GetRequestIDFromContext retrieves the request ID from the context
// Returns the request ID or empty string if not found
func GetRequestIDFromContext(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey("request_id")).(string); ok {
		return requestID
	}
	return ""
}

// setRequestIDToContext sets the request ID in the context
func setRequestIDToContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey("request_id"), requestID)
}

// generateRequestID generates a new UUID for request ID
func generateRequestID() string {
	return uuid.New().String()
}
