package helper

import (
	"log"
	"os"
	"voucher_system/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLogger(config config.Config) (*zap.Logger, error) {

	logLevel := zap.InfoLevel
	if config.AppDebug {
		logLevel = zap.DebugLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		MessageKey:     "M",
		CallerKey:      "C",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicf("Failed to open log file: %v", err)
		return nil, err
	}

	defer file.Close()

	core := zapcore.NewTee(

		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(file),
			logLevel,
		),

		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			logLevel,
		),
	)

	logger := zap.New(core)
	logger.Info("Logger initialized successfully")

	return logger, nil
}
