package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreateExtratoService(domain domain.ExtratoDomain) {
	logger.Info("Init CreateFundosService", "sincronizarFundos")

	fs.repository.CreateExtratoRepository(domain)
	logger.Info("Finish CreateFundosService", "sincronizarFundos")
}
