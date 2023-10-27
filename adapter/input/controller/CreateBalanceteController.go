package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateBalanceteController(request request.BalanceteRequest) {
	logger.Info("Init CreateBalanceteController", "sincronizarFundos")

	domain := &domain.BalanceteDomain{}
	copier.Copy(domain, request)
	fc.service.CreateBalanceteService(*domain)
	logger.Info("Finish CreateBalanceteController", "sincronizarFundos")

}
