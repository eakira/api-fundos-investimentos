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
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("About to start user application", "start")

	err := LoadDotEnv()
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
	fundosController.CreateTopicController() // Carregando os topicos do kafka

	initServicos(fundosController)

}

func LoadDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func initServicos(fundosController controller.FundosControllerInterface) {
	router := gin.Default()
	addr := env.GetPort()
	chanError := make(chan error)
	chanShutdown := make(chan bool)
	var mutex = &sync.Mutex{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", addr),
		Handler: router,
	}

	go initHttp(router, server, fundosController, mutex)
	go initListener(fundosController, chanShutdown, mutex)
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
		timeout, err := time.ParseDuration(env.GetTimeShutdown())
		if err != nil {
			logger.Error(
				"Variavel de ambiente do TIME_SHUTDOWN não está configurada corretamente",
				err,
				"shutdown",
			)
			if env.GetAmbienteDev() {
				panic(err)
			}
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
	mutex *sync.Mutex,
) {
	routes.InitRoutes(&router.RouterGroup, fundosController, mutex)

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
func initListener(
	fundosController controller.FundosControllerInterface,
	chanShutdown chan bool,
	mutex *sync.Mutex,
) {
	listener.Consume(fundosController, chanShutdown, mutex)
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
