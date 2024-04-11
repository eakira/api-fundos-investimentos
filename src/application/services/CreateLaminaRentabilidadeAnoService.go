package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateLaminaRentabilidadeAnoService(rentabilidadeAno []domain.LaminaRentabilidadeAnoDomain) *resterrors.RestErr {
	logger.Info("Init CreateLaminaRentabilidadeAnoService", "sincronizarFundos")

	if err := fs.repository.CreateManyLaminaRentabilidadeAnoRepository(rentabilidadeAno); err != nil {
		logger.Error("Error calling CreateManyLaminaRentabilidadeAnoRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateLaminaRentabilidadeAnoService", "sincronizarFundos")
	return nil
}
