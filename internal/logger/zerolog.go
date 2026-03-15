package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

// ZerologLogger implements the Logger interface using zerolog
type ZerologLogger struct {
	logger zerolog.Logger
}

// NewZerologLogger creates a new zerolog-based logger
func NewZerologLogger() Logger {
	// Configure zerolog with console writer for development
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	
	// Use console writer for pretty printing
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	
	logger := zerolog.New(output).With().Timestamp().Logger()
	
	return &ZerologLogger{
		logger: logger,
	}
}

// NewZerologLoggerWithLevel creates a new zerolog-based logger with specified level
func NewZerologLoggerWithLevel(level zerolog.Level) Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(level)
	
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	
	logger := zerolog.New(output).With().Timestamp().Logger()
	
	return &ZerologLogger{
		logger: logger,
	}
}

// NewZerologLoggerFromExisting creates a new zerolog-based logger from existing zerolog.Logger
func NewZerologLoggerFromExisting(logger zerolog.Logger) Logger {
	return &ZerologLogger{
		logger: logger,
	}
}

// Debug logs a debug message with optional fields
func (z *ZerologLogger) Debug(msg string, fields ...Field) {
	event := z.logger.Debug()
	for _, field := range fields {
		event = event.Interface(field.Key, field.Value)
	}
	event.Msg(msg)
}

// Info logs an info message with optional fields
func (z *ZerologLogger) Info(msg string, fields ...Field) {
	event := z.logger.Info()
	for _, field := range fields {
		event = event.Interface(field.Key, field.Value)
	}
	event.Msg(msg)
}

// Warn logs a warning message with optional fields
func (z *ZerologLogger) Warn(msg string, fields ...Field) {
	event := z.logger.Warn()
	for _, field := range fields {
		event = event.Interface(field.Key, field.Value)
	}
	event.Msg(msg)
}

// Error logs an error message with optional fields
func (z *ZerologLogger) Error(msg string, fields ...Field) {
	event := z.logger.Error()
	for _, field := range fields {
		event = event.Interface(field.Key, field.Value)
	}
	event.Msg(msg)
}

// Fatal logs a fatal message with optional fields and exits
func (z *ZerologLogger) Fatal(msg string, fields ...Field) {
	event := z.logger.Fatal()
	for _, field := range fields {
		event = event.Interface(field.Key, field.Value)
	}
	event.Msg(msg)
}

// With returns a new logger with the given fields added
func (z *ZerologLogger) With(fields ...Field) Logger {
	newLogger := z.logger.With()
	for _, field := range fields {
		newLogger = newLogger.Interface(field.Key, field.Value)
	}
	
	return &ZerologLogger{
		logger: newLogger.Logger(),
	}
}

// WithStr returns a new logger with the given string field added
func (z *ZerologLogger) WithStr(key, value string) Logger {
	newLogger := z.logger.With().Str(key, value).Logger()
	return &ZerologLogger{
		logger: newLogger,
	}
}

// WithInt returns a new logger with the given int field added
func (z *ZerologLogger) WithInt(key string, value int) Logger {
	newLogger := z.logger.With().Int(key, value).Logger()
	return &ZerologLogger{
		logger: newLogger,
	}
}

// WithErr returns a new logger with the given error field added
func (z *ZerologLogger) WithErr(err error) Logger {
	if err == nil {
		return z
	}
	newLogger := z.logger.With().Err(err).Logger()
	return &ZerologLogger{
		logger: newLogger,
	}
}

// WithCtx returns a new logger with context fields added
func (z *ZerologLogger) WithCtx(ctx context.Context) Logger {
	newLogger := z.logger.With()
	
	// Note: Request ID is now handled by middleware package
	// This method can be extended to add other context fields as needed
	
	return &ZerologLogger{
		logger: newLogger.Logger(),
	}
}

// Config holds logger configuration
type Config struct {
	Level  string // debug, info, warn, error, fatal
	Format string // console, json
}

// NewLoggerFromConfig creates a new logger with the given configuration
func NewLoggerFromConfig(cfg Config) Logger {
	// Parse log level
	level := parseLogLevel(cfg.Level)
	zerolog.SetGlobalLevel(level)

	// Configure time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var logger zerolog.Logger
	
	// Configure output format based on configuration
	switch cfg.Format {
	case "json":
		// JSON output for production
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	default:
		// Console output for development (default)
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	}

	return &ZerologLogger{
		logger: logger,
	}
}

// parseLogLevel converts string log level to zerolog.Level
func parseLogLevel(level string) zerolog.Level {
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

// NewLoggerFromEnv creates a new logger using environment variables
func NewLoggerFromEnv() Logger {
	cfg := Config{
		Level:  getEnvOrDefault("LOG_LEVEL", "info"),
		Format: getEnvOrDefault("LOG_FORMAT", "console"),
	}
	return NewLoggerFromConfig(cfg)
}

// getEnvOrDefault returns environment variable value or default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
