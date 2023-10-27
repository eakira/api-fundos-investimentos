package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreateFundosService(domain domain.FundosDomain) {
	logger.Info("Init CreateFundosService", "sincronizarFundos")

	fs.repository.CreateFundosRepository(domain)
	logger.Info("Finish CreateFundosService", "sincronizarFundos")
}
