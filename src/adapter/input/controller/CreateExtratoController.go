package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateExtratoController(request []request.ExtratoRequest) {
	logger.Info("Init CreateExtratoController", "sincronizarFundos")

	extratoDomain := make([]domain.ExtratoDomain, len(request))
	copier.Copy(extratoDomain, request)

	if err := fc.service.CreateExtratoService(extratoDomain); err != nil {
		logger.Error("Error calling CreateExtratoService", err, "sincronizarFundos")
	}

	logger.Info("Finish CreateExtratoController", "sincronizarFundos")
}
