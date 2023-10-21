package controller

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/core/dto"

	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) SincronizarFundos(c *gin.Context) {
	logger.Info("Init SincronizarFundos", "sincronizarFundos")
	dto := dto.FundosQueueDto{
		Topic: env.GetTopicSincronizar(),
		Queue: "update-all",
		Data:  nil,
	}

	fc.service.QueueFundosSincronizarService(dto)
	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}
