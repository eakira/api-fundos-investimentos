package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateExtratoController(request []request.ExtratoRequest) {
	logger.Info("Init CreateExtratoController", "sincronizarFundos")

	domain := &[]domain.ExtratoDomain{}
	copier.Copy(domain, request)
	fc.service.CreateExtratoService(*domain)
	logger.Info("Finish CreateExtratoController", "sincronizarFundos")

}
