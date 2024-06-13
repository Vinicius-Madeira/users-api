package repository

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		UserDomain model.UserDomainInterface,
	) (model.UserDomainInterface, rest_err.RestError)
}