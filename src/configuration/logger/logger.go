package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_PATH   = "logs/go.log"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	os.OpenFile(LOG_PATH, os.O_RDONLY|os.O_CREATE, 0666)
	logConfig := zap.Config{
		OutputPaths: []string{"stdout", LOG_PATH},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfig.Build()
}

func Info(message string, journey string, tags ...zap.Field) {
	tags = append(tags, zap.String("journey", journey))

	log.Info(message, tags...)
	errSync := log.Sync()
	if errSync != nil {
		return
	}
}

func Error(message string, err error, journey string, tags ...zap.Field) {
	tags = append(tags, zap.String("journey", journey))
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	errSync := log.Sync()
	if errSync != nil {
		return
	}
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
