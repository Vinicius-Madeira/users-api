package main

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/controller"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/routes"
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

	// Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
