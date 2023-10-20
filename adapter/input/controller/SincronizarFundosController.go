package controller

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"

	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) SincronizarFundos(c *gin.Context) {
	logger.Info("Init SincronizarFundos", "sincronizarFundos")

	fc.service.QueueFundosExternoService(env.GetTopicSincronizar(), "update-all")
	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}
