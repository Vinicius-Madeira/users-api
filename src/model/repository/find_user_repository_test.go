package repository_test

import (
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository"
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_database_test"

	_ = os.Setenv("MONGODB_COLLECTION_NAME", collectionName)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_send_a_valid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "testing123$",
			Name:     "findUserByEmail",
			Age:      50,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomain.GetPassword(), "testing123$")
		assert.EqualValues(t, userDomain.GetName(), "findUserByEmail")
		assert.EqualValues(t, userDomain.GetAge(), 50)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test@test.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByID(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_database_test"

	_ = os.Setenv("MONGODB_COLLECTION_NAME", collectionName)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_send_a_valid_id_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "testing123$",
			Name:     "findUserByEmail",
			Age:      50,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomain.GetPassword(), "testing123$")
		assert.EqualValues(t, userDomain.GetName(), "findUserByEmail")
		assert.EqualValues(t, userDomain.GetAge(), 50)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("testID")

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, fmt.Sprintf("User not found with this ID: %s", "testID"))
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_database_test"

	_ = os.Setenv("MONGODB_COLLECTION_NAME", collectionName)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_send_a_valid_email_and_password_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "testing123$",
			Name:     "findUserByEmail",
			Age:      50,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomain.GetPassword(), "testing123$")
		assert.EqualValues(t, userDomain.GetName(), "findUserByEmail")
		assert.EqualValues(t, userDomain.GetAge(), 50)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("test", "testing123$")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))
		databaseMock := mt.Client.Database(databaseName)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("test@test.com", "testing123$")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
