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
	"fmt"
	"log"
	"os"
	"strconv"

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

	fundosController := initDependenciesController(database)

	x := os.Args
	if len(x) > 1 {
		partition := 0
		if len(x) == 3 {
			partition, _ = strconv.Atoi(x[2])
			fmt.Println(int32(partition))
		}
		listener.Consume(x[1], int32(partition), fundosController)
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, fundosController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

func initDependenciesController(
	database *mongo.Database,
) controller.FundosControllerInterface {
	fundosHttp := externo.NewFundosClient()
	fundosQueue := queue.NewProduce()
	fundosRepo := repository.NewFundosRepository(database)
	userService := services.NewFundosDomainService(fundosRepo, fundosQueue, fundosHttp)
	return controller.NewFundosControllerInterface(userService)
}
