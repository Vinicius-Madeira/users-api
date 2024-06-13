package repository

import (
	"context"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"os"
)

const mongodbCollection = "MONGODB_COLLECTION_NAME"

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestError) {

	logger.Info("Init createUser repository")
	collectionName := os.Getenv(mongodbCollection)

	collection := ur.databaseConnection.Collection(collectionName)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
