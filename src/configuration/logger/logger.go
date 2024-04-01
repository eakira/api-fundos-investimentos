package logger

import (
	"api-fundos-investimentos/configuration/env"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT     = env.GetLogOutup()
	LOG_PATH       = env.GetLogInfoLevel()
	LOG_ERROR_PATH = env.GetLogErrorPath()
	LOG_LEVEL      = env.GetLogLevel()
)

func init() {
	// Abra ou crie o arquivo de log de erro
	errFile, err := os.OpenFile(LOG_ERROR_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer errFile.Close() // Feche o arquivo depois de usar

	logConfig := zap.Config{
		OutputPaths:      []string{"stdout", LOG_PATH},
		ErrorOutputPaths: []string{LOG_ERROR_PATH},
		Level:            zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:         "json",
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

	log.Error(message, tags...)
	log.Sync()
	panic(err)
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(LOG_LEVEL)) {
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
