# Logger Package

This package provides a complete logging solution with zerolog, including abstraction interfaces, middleware, and context utilities.

## Features

- **Abstracted Logger Interface**: Clean interface that abstracts zerolog implementation
- **Logger Middleware**: Attaches logger to request context with request context
- **Context Utilities**: Easy retrieval and injection of logger in context
- **Structured Logging**: JSON-formatted logs with consistent fields
- **Configuration Support**: Built-in configuration from environment variables or config struct

## Usage

### Basic Setup

```go
import "go-api-project/internal/logger"

// Create logger with environment variables (LOG_LEVEL, LOG_FORMAT)
baseLogger := logger.NewLoggerFromEnv()

// Or create with explicit config
baseLogger := logger.NewLoggerFromConfig(logger.Config{
    Level:  "info",
    Format: "console",
})
```

### In Middleware

```go
import (
    "go-api-project/internal/logger"
    "go-api-project/internal/middleware"
)

// Apply middleware (note: RequestIDMiddleware should come first)
router.Use(
    middleware.RequestIDMiddleware(),    // First: generate request ID
    logger.LoggerMiddleware(baseLogger), // Second: add logger with request ID
    middleware.RecoveryMiddleware(),    // Third: handle panics
)
```

### In Handlers

```go
func (h *UserHandler) CreateUser(c *gin.Context) {
    // Get logger from context
    log := logger.GetLoggerFromContext(c.Request.Context())
    
    log.Info("Creating user")
    
    // Add context-specific fields
    userLogger := log.WithStr("user_id", userID).WithStr("action", "create")
    
    // Log errors
    if err := h.userService.CreateUser(user); err != nil {
        userLogger.WithErr(err).Error("Failed to create user")
        c.JSON(500, gin.H{"error": "Internal server error"})
        return
    }
    
    userLogger.Info("User created successfully")
    c.JSON(201, gin.H{"message": "User created"})
}
```

## Logger Interface Methods

```go
type Logger interface {
    Debug(msg string, fields ...Field)
    Info(msg string, fields ...Field)
    Warn(msg string, fields ...Field)
    Error(msg string, fields ...Field)
    Fatal(msg string, fields ...Field)
    With(fields ...Field) Logger
    WithStr(key, value string) Logger
    WithInt(key string, value int) Logger
    WithErr(err error) Logger
    WithCtx(ctx context.Context) Logger
}
```

### Field Creation

```go
// Create different types of fields
logger.Info("User action", 
    logger.Str("user_id", "123"),
    logger.Int("attempts", 3),
    logger.Err(someError),
    logger.Any("metadata", customData),
)

// Chain field methods
logger.WithStr("user_id", "123").
    WithInt("attempts", 3).
    WithErr(someError).
    Info("User action completed")
```

### Context Utilities

```go
// Get logger from context (returns default logger if not found)
log := logger.GetLoggerFromContext(ctx)

// Set logger in context
ctx = logger.SetLoggerToContext(ctx, customLogger)

// Get request ID from context
requestID := logger.GetRequestIDFromContext(ctx)
```

## Configuration

### Environment Variables

```bash
LOG_LEVEL=info          # debug, info, warn, error, fatal
LOG_FORMAT=console      # console (for development) or json (for production)
```

### Configuration Struct

```go
cfg := logger.Config{
    Level:  "info",
    Format: "console",
}
baseLogger := logger.NewLoggerFromConfig(cfg)
```

## Log Output

The logger outputs structured JSON logs with the following automatic fields:

- `timestamp`: Request timestamp
- `level`: Log level (debug, info, warn, error, fatal)
- `request_id`: Unique request identifier (added by middleware)
- `method`: HTTP method
- `path`: Request path
- `remote_addr`: Client IP address
- `user_agent`: User agent string
- `duration`: Request duration (for completion logs)
- `status_code`: HTTP status code (for completion logs)

Example log output:
```json
{
  "level": "info",
  "timestamp": 1672531200,
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "method": "POST",
  "path": "/api/v1/users",
  "remote_addr": "192.168.1.100",
  "user_agent": "Mozilla/5.0...",
  "user_id": "12345",
  "action": "create",
  "message": "User created successfully"
}
```

## Files Structure

- `interfaces.go` - Logger interface and field types
- `zerolog.go` - Zerolog implementation and configuration
- `context.go` - Context utilities and constants
- `middleware.go` - Logger middleware for Gin
- `README.md` - This documentation

## Integration with Middleware Package

The logger package is designed to work seamlessly with the middleware package:

```go
import (
    "go-api-project/internal/logger"
    "go-api-project/internal/middleware"
)

// Recommended middleware order
router.Use(
    middleware.RequestIDMiddleware(),    // Generates request ID
    logger.LoggerMiddleware(baseLogger), // Creates logger with request ID
    middleware.RecoveryMiddleware(),    // Handles panics
)
```

## Testing

For testing, you can create a mock logger:

```go
type MockLogger struct {
    logs []string
}

func (m *MockLogger) Info(msg string, fields ...logger.Field) {
    m.logs = append(m.logs, msg)
}

// ... implement other methods

// Use in tests
mockLogger := &MockLogger{}
ctx := logger.SetLoggerToContext(context.Background(), mockLogger)
```
