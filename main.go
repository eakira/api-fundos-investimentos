package main

import (
	"api-fundos-investimentos/adapter/controller"
	"api-fundos-investimentos/adapter/controller/routes"
	"api-fundos-investimentos/adapter/output/repository"
	"api-fundos-investimentos/application/services"
	"api-fundos-investimentos/configuration/database/mongodb"
	"api-fundos-investimentos/configuration/logger"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("About to start user application", "start")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

func initDependencies(
	database *mongo.Database,
) controller.FundosControllerInterface {
	fundosRepo := repository.NewFundosRepository(database)
	userService := services.NewUserDomainService(fundosRepo)
	return controller.NewFundosControllerInterface(userService)
}
