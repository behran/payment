package logger

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger ...
var Logger *zap.Logger

// New ...
func New() *zap.Logger {
	var err error

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	Logger, err = config.Build()
	if err != nil {
		log.Fatal(err)
	}
	return Logger
}
