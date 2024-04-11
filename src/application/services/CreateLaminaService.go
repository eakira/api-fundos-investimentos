package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreateLaminaService(domain []domain.LaminaDomain) {
	logger.Info("Init CreateLaminaService", "sincronizarFundos")

	fs.repository.CreateManyLaminaRepository(domain)
	logger.Info("Finish CreateLaminaService", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateLaminaCarteiraService(domain []domain.LaminaCarteiraDomain) {
	logger.Info("Init CreateLaminaCarteiraService", "sincronizarFundos")

	fs.repository.CreateManyLaminaCarteiraRepository(domain)
	logger.Info("Finish CreateLaminaCarteiraService", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateLaminaRentabilidadeAnoService(domain []domain.LaminaRentabilidadeAnoDomain) {
	logger.Info("Init CreateLaminaRentabilidadeAnoService", "sincronizarFundos")

	fs.repository.CreateManyLaminaRentabilidadeAnoRepository(domain)
	logger.Info("Finish CreateLaminaRentabilidadeAnoService", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateLaminaRentabilidadeMesService(domain []domain.LaminaRentabilidadeMesDomain) {
	logger.Info("Init CreateLaminaRentabilidadeMesService", "sincronizarFundos")

	fs.repository.CreateManyLaminaRentabilidadeMesRepository(domain)
	logger.Info("Finish CreateLaminaRentabilidadeMesService", "sincronizarFundos")
}
