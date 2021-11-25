// Package logger wraps around the zap.Logger for convenience and to ensure we use singleton instance of the logger
package logger

import (
	"context"
	_log "log"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/Tommy647/go_example/internal/trace"
)

var (
	// l singleton logger instance can be accessed through the wrapper functions (Debug, Info, Fatal)
	l *zap.Logger
	// applicationName for logging
	applicationName string
)

const loggerNotInitialised = `logger not initialised`
const wrapperStackOffset = 2 // I <3 magic numbers :/

// New instance of the logger, runs as a singleton
// so subsequent calls will return the same logger
func New(appName string) error {
	applicationName = appName
	if l == nil {
		// get a default (sane) config
		config := zap.NewProductionConfig()
		// set the minimum log level
		config.Level = getLevel()
		// create our logger
		logger, err := config.Build(
			zap.WithCaller(true),                                        // prints the line log was called from, helps debugging
			zap.AddCallerSkip(wrapperStackOffset),                       // adds an offset so caller is the calling code, not the wrapper
			zap.Fields(zap.String("application_name", applicationName)), // default field to all logs
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
func Close() error { return l.Sync() }

// Debug level message
func Debug(ctx context.Context, m string, fields ...zap.Field) {
	if l == nil {
		_log.Println(loggerNotInitialised)
		return
	}
	fields = addDefaultFields(ctx, fields)
	l.Debug(m, fields...)
	// ok this is tedious, but having these wrapper functions
	// makes the custom logger easier to use and consistent
}

// Info level message
func Info(ctx context.Context, m string, fields ...zap.Field) {
	log(ctx, m, l.Info, fields...) // this is a bit cleaner
}

// Warn level message
func Warn(ctx context.Context, m string, fields ...zap.Field) { log(ctx, m, l.Warn, fields...) }

// Error level message
func Error(ctx context.Context, m string, fields ...zap.Field) { log(ctx, m, l.Error, fields...) }

// Fatal level message
func Fatal(ctx context.Context, m string, fields ...zap.Field) { log(ctx, m, l.Fatal, fields...) }

// log generic log caller, so we are consistent without fields and this lib remains easy to use
func log(ctx context.Context, m string, f func(string, ...zap.Field), fields ...zap.Field) {
	// check we have a logger
	if l == nil {
		_log.Println(loggerNotInitialised)
	}
	fields = addDefaultFields(ctx, fields)
	f(m, fields...)
}

// addDefaultFields to each log line
func addDefaultFields(ctx context.Context, fields []zap.Field) []zap.Field {
	return append(fields,
		zap.String("trace_id", trace.GetTraceID(ctx)), // trace is dynamic to each request
	)
}

// get the logging level from the env vars, log lines 'below'
// this level will be ignored and not logged
func getLevel() zap.AtomicLevel {
	switch strings.ToLower(os.Getenv(`log_level`)) {
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case `debug`:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel) // info by default
	}
}
