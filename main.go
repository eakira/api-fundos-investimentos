package main

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/adapter/input/controller/routes"
	"api-fundos-investimentos/adapter/input/listener"
	"api-fundos-investimentos/adapter/output/externo"
	"api-fundos-investimentos/adapter/output/queue"
	"api-fundos-investimentos/adapter/output/repository"
	"api-fundos-investimentos/application/services"
	"api-fundos-investimentos/configuration/database/mongodb"
	"api-fundos-investimentos/configuration/logger"
	"context"
	"log"
	"os"

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

	x := os.Args
	if len(x) > 1 {
		listener.Consume(x[1])
	}

	fundosController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, fundosController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

func initDependencies(
	database *mongo.Database,
) controller.FundosControllerInterface {
	fundosHttp := externo.NewFundosClient()
	fundosQueue := queue.NewProduce()
	fundosRepo := repository.NewFundosRepository(database)
	userService := services.NewFundosDomainService(fundosRepo, fundosQueue, fundosHttp)
	return controller.NewFundosControllerInterface(userService)
}
