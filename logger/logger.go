package logger

import (
	"os"

	"go.uber.org/zap"
)

var zapLogger *zap.SugaredLogger

func InitLogger() {
	if zapLogger != nil {
		return
	}
	logger, _ := zap.NewDevelopment()

	env := os.Getenv("GIN_MODE")
	if env != "" && env == "release" {
		logger, _ = zap.NewProduction()
	}
	zapLogger = logger.Sugar()
}

func Debugf(template string, args ...interface{}) {
	zapLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	zapLogger.Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	zapLogger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	zapLogger.Fatalf(template, args...)
}

func Sync() {
	_ = zapLogger.Sync()
}
