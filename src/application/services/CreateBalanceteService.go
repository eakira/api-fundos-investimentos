package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreateBalanceteService(domain []domain.BalanceteDomain) {
	logger.Info("Init CreateFundosService", "sincronizarFundos")

	fs.repository.CreateManyBalecenteRepository(domain)
	logger.Info("Finish CreateFundosService", "sincronizarFundos")
}
