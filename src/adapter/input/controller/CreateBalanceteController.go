package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateBalanceteController(request []request.BalanceteRequest) {
	logger.Info("Init CreateBalanceteController", "sincronizarFundos")

	balanceteDomain := make([]domain.BalanceteDomain, len(request))
	copier.Copy(balanceteDomain, request)

	if err := fc.service.CreateBalanceteService(balanceteDomain); err != nil {
		logger.Error("Error calling CreateBalanceteService", err, "sincronizarFundos")
	}

	logger.Info("Finish CreateBalanceteController", "sincronizarFundos")
}
