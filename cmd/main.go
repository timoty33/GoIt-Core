package main

import (
	"errors"

	"github.com/timoty33/goit-core/internal/infra/mongo"
	"github.com/timoty33/goit-core/pkg/config"
	"github.com/timoty33/goit-core/pkg/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// init logger
	logger.Create()
	defer logger.Sync()
	logger.Log(
		zapcore.InfoLevel,
		"Logger initialized",
		zap.String("locate", "cmd/main.go"),
	)

	// load .env
	logger.Log(
		zapcore.InfoLevel,
		"Loading .env",
		zap.String("locate", "cmd/main.go"),
	)
	err := config.LoadEnv()
	if err != nil {
		panic(err)
	}
	logger.Log(
		zapcore.InfoLevel,
		".env loaded",
		zap.String("locate", "cmd/main.go"),
	)

	logger.Log(
		zapcore.InfoLevel,
		"Conecting MongoDB",
		zap.String("locate", "cmd/main.go"),
	)
	errorMongo := mongo.Connect()
	if errorMongo.Err != nil {
		logger.Log(
			zap.ErrorLevel, "Error when conecting MongoDB",
			zap.String("locate", "cmd/main.go"),
			zap.Error(errors.New(errorMongo.Error())),
		)

		panic(errorMongo.Error())
	}
	logger.Log(
		zapcore.InfoLevel,
		"MongoDB connected",
		zap.String("locate", "cmd/main.go"),
	)

	defer mongo.Disconnect()
}
