package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateLaminaRentabilidadeMesService(rentabilidadeMes []domain.LaminaRentabilidadeMesDomain) *resterrors.RestErr {
	logger.Info("Init CreateLaminaRentabilidadeMesService", "sincronizarFundos")

	if err := fs.repository.CreateManyLaminaRentabilidadeMesRepository(rentabilidadeMes); err != nil {
		logger.Error("Error calling CreateManyLaminaRentabilidadeMesRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateLaminaRentabilidadeMesService", "sincronizarFundos")
	return nil
}
