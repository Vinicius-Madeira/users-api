package repository

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const mongodbCollection = "MONGODB_COLLECTION_NAME"

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		UserDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestError)
	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestError)
	FindUserByEmailAndPassword(
		email string,
		password string,
	) (model.UserDomainInterface, *rest_err.RestError)
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestError)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestError
	DeleteUser(userId string) *rest_err.RestError
}
