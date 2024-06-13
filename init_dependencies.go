package main

import (
	"github.com/Vinicius-Madeira/go-web-app/src/controller"
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository"
	"github.com/Vinicius-Madeira/go-web-app/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) (controller.UserControllerInterface, error) {

	repo := repository.NewUserRepository(database)
	serv := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(serv), nil
}
