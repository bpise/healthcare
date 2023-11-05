package logger

import (
	"os"

	"go.uber.org/zap"
)

// zapLogger holds the logger instance for the entire package.
var zapLogger *zap.SugaredLogger

// InitLogger initializes the logger based on the runtime environment.
func InitLogger() {
	// Check if the logger has already been initialized.
	if zapLogger != nil {
		return
	}

	// Create a new logger instance for development mode.
	logger, _ := zap.NewDevelopment()

	// Check the GIN_MODE environment variable to switch to production mode if set as release.
	env := os.Getenv("GIN_MODE")
	if env != "" && env == "release" {
		logger, _ = zap.NewProduction()
	}
	zapLogger = logger.Sugar()
}

// Debugf logs a message at the debug log level.
func Debugf(template string, args ...interface{}) {
	zapLogger.Debugf(template, args...)
}

// Infof logs a message at the info log level.
func Infof(template string, args ...interface{}) {
	zapLogger.Infof(template, args...)
}

// Errorf logs a message at the error log level.
func Errorf(template string, args ...interface{}) {
	zapLogger.Errorf(template, args...)
}

// Fatalf logs a message at the fatal log level.
func Fatalf(template string, args ...interface{}) {
	zapLogger.Fatalf(template, args...)
}

func Sync() {
	_ = zapLogger.Sync()
}
