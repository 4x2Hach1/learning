package utils

import (
	"log"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func Logger() *zap.Logger {
	if logger == nil {
		zapLogger, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}

		logger = zapLogger
		return logger
	} else {
		return logger
	}
}

func LoggingInfo() func(string, ...interface{}) {
	return Logger().Sugar().Infow
}

func LoggingWarn() func(string, ...interface{}) {
	return Logger().Sugar().Warnw
}

func LoggingFatal() func(string, ...interface{}) {
	return Logger().Sugar().Fatalw
}
