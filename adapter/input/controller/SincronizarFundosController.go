package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"

	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) SincronizarFundosController(c *gin.Context) {
	logger.Info("Init SincronizarFundos", "sincronizarFundos")
	dto := request.FundosQueueRequest{
		Topic: env.GetTopicSincronizar(),
		Queue: "update-all",
		Data:  nil,
	}

	fc.service.QueueFundosSincronizarService(dto)
	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}
