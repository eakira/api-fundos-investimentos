package controller

import (
	"api-fundos-investimentos/configuration/logger"

	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) SincronizarFundosController(c *gin.Context) {
	logger.Info("Init SincronizarFundos", "sincronizarFundos")

	fc.service.QueueFundosSincronizarService("all")

	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}
