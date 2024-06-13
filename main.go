package main

import (
	"context"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/database/mongodb"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/controller"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/routes"
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository"
	"github.com/Vinicius-Madeira/go-web-app/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("Starting application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database, error=%s\n", err.Error())
		return
	}
	// Init dependencies
	repo := repository.NewUserRepository(database)
	serv := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(serv)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
