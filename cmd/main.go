package main

import (
	"github.com/timoty33/goit-core/internal/infra/mongo"
	"github.com/timoty33/goit-core/pkg/config"
	"github.com/timoty33/goit-core/pkg/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// carrega o .env
	err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	// inicia o logger
	logger.Create()
	defer logger.Sync()
	logger.Log(zapcore.InfoLevel, "Logger iniciado",
		zap.String("locate", "main.go"))

	// conecta no MongoDB
	err = mongo.Connect()
	if err != nil {
		logger.Log(zapcore.ErrorLevel, "Erro ao conectar no MongoDB",
			zap.String("locate", "main.go"),
			zap.Error(err))
		panic(err)
	}

	defer mongo.Disconnect()
}
