package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

// Connect abre a conexão com o MongoDB
func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return fmt.Errorf("MONGO_URI não definido")
	}

	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		return fmt.Errorf("MONGO_DB não definido")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("erro ao conectar no MongoDB: %w", err)
	}

	// Testa a conexão
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("erro ao fazer ping no MongoDB: %w", err)
	}

	Client = client
	DB = client.Database(dbName)

	fmt.Println("✅ Conectado ao MongoDB:", dbName)
	return nil
}

// Disconnect fecha a conexão (chamado no shutdown)
func Disconnect() error {
	if Client == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return Client.Disconnect(ctx)
}
