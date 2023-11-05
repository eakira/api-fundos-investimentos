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
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	TIME_SHUTDOWN = "TIME_SHUTDOWN"
	PORT          = "PORT"
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

	router := gin.Default()
	addr := os.Getenv(PORT)
	chanError := make(chan error)
	chanShutdown := make(chan bool)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", addr),
		Handler: router,
	}

	go initHttp(router, server, fundosController)
	go initListener(fundosController, chanShutdown)
	go GraceFullyShutdown(server, chanError, chanShutdown, fundosController)

	if err := <-chanError; err != nil {
		logger.Error("Error: ", err, "init")
	}

}
func GraceFullyShutdown(
	server *http.Server,
	chanError chan error,
	chanShutdown chan bool,
	fundosController controller.FundosControllerInterface,
) {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()
		logger.Info("Recebido sinal para desligar...", "shutdown")
		timeout, err := time.ParseDuration(os.Getenv(TIME_SHUTDOWN))
		if err != nil {
			logger.Error(
				"Variavel de ambiente do TIME_SHUTDOWN não está configurada corretamente",
				err,
				"shutdown",
			)
			panic(err)
		}
		ctxTimeout, cancel := context.WithTimeout(context.Background(), timeout)

		defer func() {
			stop()
			cancel()
			close(chanError)
		}()
		err = server.Shutdown(ctxTimeout)
		if status := <-chanShutdown; status {
			logger.Error("Desligando o listener...", err, "shutdown")
			close(chanShutdown)
		}

		if err != nil {
			logger.Error("Erro no shutdown", err, "shutdown")
		}
	}()

}
func initHttp(
	router *gin.Engine,
	server *http.Server,
	fundosController controller.FundosControllerInterface,
) {
	routes.InitRoutes(&router.RouterGroup, fundosController)

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
func initListener(
	fundosController controller.FundosControllerInterface,
	chanShutdown chan bool,
) {
	partition := 0
	chanShutdown <- true
	listener.Consume(int32(partition), fundosController, chanShutdown)
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
