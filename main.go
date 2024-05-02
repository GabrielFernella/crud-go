package main

import (
	"context"
	"log"

	"github.com/GabrielFernella/crud-go/src/configuration/database/mongodb"
	"github.com/GabrielFernella/crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error(),
		)
		return
	}

	// Init dependencies
	userController := initDependencies(database)

	// repo := repository.NewUserRepository(database)
	// service := service.NewUserDomainService(repo)
	// userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
