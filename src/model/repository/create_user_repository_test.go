package repository_test

import (
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_database_test"

	_ = os.Setenv("MONGODB_COLLECTION_NAME", collectionName)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.CreateUser(model.NewUserDomain(
			"test@test.com",
			"testing123$",
			"unitTestRepository",
			36,
		))
		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomain.GetPassword(), "testing123$")
		assert.EqualValues(t, userDomain.GetName(), "unitTestRepository")
		assert.EqualValues(t, userDomain.GetAge(), 36)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@test.com",
			"testing123$",
			"unitTestRepository",
			36,
		)
		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
