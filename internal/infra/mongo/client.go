package mongo

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	e "github.com/timoty33/goit-core/internal/domain/errors"
	"github.com/timoty33/goit-core/pkg/logger"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

// Connect abre a conexão com o MongoDB
func Connect() e.GoItError {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		logger.Log(
			zapcore.PanicLevel,
			"Can't get MONGO_URI from .env",
			zap.String("locate", "internal/infra/mongo/client.go"),
		)

		return e.New("ERROR DB", "Can't get MONGO_URI from .env", errors.New("can't get MONGO_URI from .env"))
	}

	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		logger.Log(
			zapcore.PanicLevel,
			"Can't get MONGO_DB from .env",
			zap.String("locate", "internal/infra/mongo/client.go"),
		)

		return e.New("ERROR DB", "Can't get MONGO_DB from .env", errors.New("can't get MONGO_DB from .env"))
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Log(
			zapcore.FatalLevel,
			"Connection with Mongo failled",
			zap.String("locate", "internal/infra/mongo/client.go"),
			zap.Error(err),
		)

		return e.New("ERROR DB", "Can't connect in Mongo", err)
	}

	// Testa a conexão
	if err := client.Ping(ctx, nil); err != nil {
		logger.Log(
			zapcore.WarnLevel,
			"Connection with Mongo failled",
			zap.String("locate", "internal/infra/mongo/client.go"),
			zap.Error(err),
		)

		return e.New("ERROR DB", "Connection with Mongo failled", err)
	}

	Client = client
	DB = client.Database(dbName)

	logger.Log(
		zapcore.InfoLevel,
		"Mongo connected",
		zap.String("locate", "internal/infra/mongo/client.go"),
	)

	return e.New("STATUS DB OK", "Mongo connected", nil)
}

// Disconnect fecha a conexão (chamado no shutdown)
func Disconnect() e.GoItError {
	if Client == nil {
		return e.New("ERROR DB", "Can't find Client", errors.New("client doesn't exist"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := Client.Disconnect(ctx)
	if err != nil {
		logger.Log(
			zapcore.WarnLevel,
			"Can't disconnect Mongo",
			zap.String("locate", "internal/infra/mongo/client.go"),
		)
		return e.New("ERROR DB", "Can't disconnect Mongo", err)
	}

	logger.Log(
		zapcore.InfoLevel,
		"Mongo disconnected",
		zap.String("locate", "internal/infra/mongo/client.go"),
	)
	return e.New("STATUS DB OK", "Mongo disconnected", nil)
}
