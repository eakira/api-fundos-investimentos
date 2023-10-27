package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreateInformacaoDiariaService(domain domain.InformacaoDiariaDomain) {
	logger.Info("Init CreateFundosService", "sincronizarFundos")

	fs.repository.CreateInformacaoDiariaRepository(domain)
	logger.Info("Finish CreateFundosService", "sincronizarFundos")
}
