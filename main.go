package main

import (
	"context"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/database/mongodb"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("Starting application")
	err := godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database, error=%s\n", err.Error())
		return
	}
	userController, _ := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
