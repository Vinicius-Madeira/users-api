package mongodb

import (
	"context"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	URL          = "MONGODB_URL"
	DatabaseName = "MONGODB_DATABASE_NAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbUri := os.Getenv(URL)
	mongodbName := os.Getenv(DatabaseName)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB!")
	return client.Database(mongodbName), nil
}
