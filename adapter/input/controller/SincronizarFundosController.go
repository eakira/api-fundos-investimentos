package controller

import (
	"api-fundos-investimentos/configuration/logger"

	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) SincronizarFundos(c *gin.Context) {
	logger.Info("Init SincronizarFundos", "sincrinizarFundos")

	fc.service.GetFundosExternoService()
	logger.Info("Finish SincronizarFundos", "sincrinizarFundos")
}
