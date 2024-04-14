package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateFundosController(request []request.FundosRequest) {
	logger.Info("Init CreateFundosController", "sincronizarFundos")

	fundoDomain := make([]domain.FundosDomain, len(request))
	copier.Copy(fundoDomain, request)

	if err := fc.service.CreateFundosService(fundoDomain); err != nil {
		logger.Error("Error calling CreateFundosService", err, "sincronizarFundos")
	}

	logger.Info("Finish CreateFundosController", "sincronizarFundos")
}
