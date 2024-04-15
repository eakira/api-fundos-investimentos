package controller

import (
	"api-fundos-investimentos/configuration/logger"

	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) SincronizarFundosController(c *gin.Context) {
	logger.Info("Init SincronizarFundos", "sincronizarFundos")

	tipo := []string{
		//		"cadastros",
		//		"balancete",
		//		"cda",
		//		"extrato",
		//		"informacao-diaria",
		//		"informacoes-complementares",
		//"lamina",
		"perfil-mensal",
	}
	for _, value := range tipo {
		fc.service.QueueFundosSincronizarService(value)
	}

	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}
