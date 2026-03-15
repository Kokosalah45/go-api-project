package logger

import (
	"context"
)

// loggerKey is the key used to store the logger in the context
type loggerKey string

// loggerKey constant for context operations
const LoggerKey loggerKey = "logger"

// setLoggerToContext sets the logger in the context
func SetLoggerToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, logger)
}

// getLoggerFromContext retrieves the logger from the context
// Returns the logger from context or a default logger if not found
func GetLoggerFromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(LoggerKey).(Logger); ok {
		return logger
	}
	// Return a default logger if none is found in context
	return NewZerologLogger()
}
