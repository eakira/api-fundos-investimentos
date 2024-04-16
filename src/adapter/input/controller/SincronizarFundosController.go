package controller

import (
	"api-fundos-investimentos/configuration/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) SincronizarFundosController(c *gin.Context) {
	logger.Info("Init SincronizarFundos", "sincronizarFundos")

	ignorar, err := strconv.ParseBool(c.PostForm("ignorar_download"))
	if err != nil {
		ignorar = false
	}
	baixar := !ignorar

	tipo := []string{
		"cadastros",
		"balancete",
		"cda",
		"extrato",
		"informacao-diaria",
		"informacoes-complementares",
		"lamina",
		"perfil-mensal",
	}
	for _, value := range tipo {
		fc.service.QueueFundosSincronizarService(value, baixar)
	}

	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}
