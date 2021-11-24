// Package logger wraps around the zap.Logger for convenience and to ensure we use singleton instance of the logger
package logger

import (
	"context"

	"go.uber.org/zap"
)

// l singleton logger instance can be accessed through the wrapper functions (Debug, Info, Fatal)
var l *zap.Logger

// New instance of the logger, runs as a singleton
// so subsequent calls will return the same logger
func New() error {
	// @todo: get logger level from env and implement
	if l == nil {
		logger, err := zap.NewProduction(
			zap.WithCaller(true),
		)
		if err != nil {
			return err
		}
		l = logger
	}
	return nil
}

// Close syncs the logger to ensure all buffered logs
// message are processed before we close the application
func Close() error {
	return l.Sync()
}

// Debug level message
func Debug(ctx context.Context, m string, fields ...zap.Field) {
	l.Debug(m, fields...)
}

// Info level message
func Info(ctx context.Context, m string, fields ...zap.Field) {
	l.Info(m, fields...)
}

// Warn level message
func Warn(ctx context.Context, m string, fields ...zap.Field) {
	l.Warn(m, fields...)
}

// Error level message
func Error(ctx context.Context, m string, fields ...zap.Field) {
	l.Error(m, fields...)
}

// Fatal level message
func Fatal(ctx context.Context, m string, fields ...zap.Field) {
	l.Fatal(m, fields...)
}
