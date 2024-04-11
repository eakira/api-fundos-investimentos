package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateInformacaoDiariaController(
	request []request.InformacaoDiariaRequest,
) {
	logger.Info("Init CreateInformacaoDiariaController", "sincronizarFundos")

	informacaoDiariaDomain := make([]domain.InformacaoDiariaDomain, len(request))
	copier.Copy(informacaoDiariaDomain, request)

	if err := fc.service.CreateInformacaoDiariaService(informacaoDiariaDomain); err != nil {
		logger.Error("Error calling CreateInformacaoDiariaService", err, "sincronizarFundos")
	}

	logger.Info("Finish CreateInformacaoDiariaController", "sincronizarFundos")

}
