package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"

	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) SincronizarFundosController(c *gin.Context) {
	logger.Info("Init SincronizarFundos", "sincronizarFundos")

	folders := []string{
		"FI",
		"FIDC",
		"FIE",
		"FII",
		"FIP",
	}

	request := request.FundosQueueRequest{
		Topic: env.GetTopicSincronizar(),
		Queue: "update-all",
		Data:  folders,
	}

	fc.service.QueueFundosSincronizarService(request)
	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}
