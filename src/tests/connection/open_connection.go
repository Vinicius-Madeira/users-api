package connection

import (
	"context"
	"fmt"
	"github.com/ory/dockertest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func OpenConnection() (database *mongo.Database, close func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Println("Error trying to open connection")
		return
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	})
	if err != nil {
		log.Fatalf("Could not create mongo container: %s", err.Error())
		return
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(
		fmt.Sprintf("mongodb://127.0.0.1:%v", resource.GetPort("27017/tcp"))))
	if err != nil {
		log.Println("Error trying to open connection")
		return
	}

	database = client.Database(os.Getenv("MONGODB_COLLECTION_NAME"))
	close = func() {
		closeErr := resource.Close()
		if closeErr != nil {
			log.Println("Error trying to close resource")
			return
		}
	}

	return
}
