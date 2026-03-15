# Middleware Package

This package contains HTTP middleware components that are independent of the logging system.

## Components

### Request ID Middleware (`requestid.go`)

Generates and tracks unique request IDs across the application.

**Features:**
- Generates UUID for each request if not present in `X-Request-ID` header
- Preserves upstream request IDs from external services
- Sets request ID in both context and response headers
- Thread-safe and performant

**Usage:**
```go
import "go-api-project/internal/middleware"

router.Use(middleware.RequestIDMiddleware())
```

### Recovery Middleware (`recovery.go`)

Recovers from panics and returns structured JSON responses.

**Features:**
- Catches panics and prevents server crashes
- Returns structured JSON error responses
- Includes request ID in error responses
- Handles different panic types (error, string, other)

**Usage:**
```go
import "go-api-project/internal/middleware"

router.Use(middleware.RecoveryMiddleware())
```

### Context Utilities (`context.go`)

Provides context management functions for request IDs.

**Functions:**
- `GetRequestIDFromContext(ctx) string` - Retrieves request ID from context
- `setRequestIDToContext(ctx, id) context` - Sets request ID in context (internal)
- `generateRequestID() string` - Generates new UUID (internal)

## Integration with Logger Package

The middleware package is designed to work seamlessly with the logger package:

```go
import (
    "go-api-project/internal/logger"
    "go-api-project/internal/middleware"
)

// Recommended middleware order
router.Use(
    middleware.RequestIDMiddleware(),    // First: generate request ID
    logger.LoggerMiddleware(baseLogger), // Second: add logger with request ID
    middleware.RecoveryMiddleware(),    // Third: handle panics
)
```

## Request Flow

1. **Request ID Middleware** generates/validates request ID and stores in context
2. **Logger Middleware** reads request ID from context and creates logger with it
3. **Your Handlers** can access both logger and request ID from context
4. **Recovery Middleware** handles any panics and includes request ID in error response

## Error Response Format

When a panic occurs, the recovery middleware returns:

```json
{
  "error": "Internal Server Error",
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "panic_error": "specific panic message"
}
```

## Best Practices

1. **Order Matters**: Always apply RequestIDMiddleware before LoggerMiddleware
2. **Upstream Compatibility**: The middleware preserves existing `X-Request-ID` headers
3. **Context Access**: Use `middleware.GetRequestIDFromContext()` to access request ID
4. **Error Handling**: Recovery middleware ensures your app never crashes from panics

## Dependencies

- `github.com/gin-gonic/gin` - HTTP framework
- `github.com/google/uuid` - UUID generation for request IDs

## Thread Safety

All middleware functions are thread-safe and can be safely used in concurrent environments.
