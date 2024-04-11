package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreatePerfilMensalService(domain []domain.PerfilMensalDomain) {
	logger.Info("Init CreatePerfilMensalService", "sincronizarFundos")

	fs.repository.CreateManyPerfilMensalRepository(domain)
	logger.Info("Finish CreatePerfilMensalService", "sincronizarFundos")
}
