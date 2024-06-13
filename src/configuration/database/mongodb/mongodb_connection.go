package mongodb

import (
	"context"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	MongodbUrl          = os.Getenv("MONGO_URL")
	MongodbDatabaseName = os.Getenv("MONGO_DATABASE_NAME")
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbUri := os.Getenv(MongodbUrl)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB!")
	return client.Database(MongodbDatabaseName), nil
}
