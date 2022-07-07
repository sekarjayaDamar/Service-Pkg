package utilities

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	Logger, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Logger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}
