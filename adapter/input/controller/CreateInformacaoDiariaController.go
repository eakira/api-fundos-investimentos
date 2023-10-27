package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateInformacaoDiariaController(
	request request.InformacaoDiariaRequest,
) {
	logger.Info("Init CreateInformacaoDiariaController", "sincronizarFundos")

	domain := &domain.InformacaoDiariaDomain{}
	copier.Copy(domain, request)
	fc.service.CreateInformacaoDiariaService(*domain)
	logger.Info("Finish CreateInformacaoDiariaController", "sincronizarFundos")

}
