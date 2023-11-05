package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateFundosController(request []request.FundosCadastrosRequest) {
	logger.Info("Init CreateFundosController", "sincronizarFundos")

	domain := &[]domain.FundosDomain{}
	copier.Copy(domain, request)
	fc.service.CreateFundosService(*domain)

	logger.Info("Finish CreateFundosController", "sincronizarFundos")
}
