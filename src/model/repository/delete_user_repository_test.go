package repository_test

import (
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_database_test"

	_ = os.Setenv("MONGODB_COLLECTION_NAME", collectionName)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_userId_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		err := repo.DeleteUser("test")

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		err := repo.DeleteUser("test")

		assert.NotNil(t, err)
	})
}
