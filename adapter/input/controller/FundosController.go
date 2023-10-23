package controller

import (
	"api-fundos-investimentos/application/port/input"

	"github.com/gin-gonic/gin"
)

func NewFundosControllerInterface(
	serviceInterface input.FundosDomainService,
) FundosControllerInterface {
	return &fundosControllerInterface{
		service: serviceInterface,
	}
}

type FundosControllerInterface interface {
	SincronizarFundosController(c *gin.Context)
	DownloadArquivosCVMController(folder string)
}

type fundosControllerInterface struct {
	service input.FundosDomainService
}
