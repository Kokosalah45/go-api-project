package logger

import "context"

// Logger defines the interface for logging operations
type Logger interface {
	// Debug logs a debug message with optional fields
	Debug(msg string, fields ...Field)
	
	// Info logs an info message with optional fields
	Info(msg string, fields ...Field)
	
	// Warn logs a warning message with optional fields
	Warn(msg string, fields ...Field)
	
	// Error logs an error message with optional fields
	Error(msg string, fields ...Field)
	
	// Fatal logs a fatal message with optional fields and exits
	Fatal(msg string, fields ...Field)
	
	// With returns a new logger with the given fields added
	With(fields ...Field) Logger
	
	// WithStr returns a new logger with the given string field added
	WithStr(key, value string) Logger
	
	// WithInt returns a new logger with the given int field added
	WithInt(key string, value int) Logger
	
	// WithErr returns a new logger with the given error field added
	WithErr(err error) Logger
	
	// WithCtx returns a new logger with context fields added
	WithCtx(ctx context.Context) Logger
}

// Field represents a log field
type Field struct {
	Key   string
	Value interface{}
}

// Str creates a string field
func Str(key, value string) Field {
	return Field{Key: key, Value: value}
}

// Int creates an int field
func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Err creates an error field
func Err(err error) Field {
	return Field{Key: "error", Value: err}
}

// Any creates a field with any value
func Any(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}
