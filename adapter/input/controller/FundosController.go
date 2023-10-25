package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
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
	DownloadArquivosCVMController(request.FundosDownloadCvmFilesQueueRequest)
}

type fundosControllerInterface struct {
	service input.FundosDomainService
}
