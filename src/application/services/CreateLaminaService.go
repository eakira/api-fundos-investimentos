package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateLaminaService(lamina []domain.LaminaDomain) *resterrors.RestErr {
	logger.Info("Init CreateLaminaService", "sincronizarFundos")

	if err := fs.repository.CreateManyLaminaRepository(lamina); err != nil {
		logger.Error("Error calling CreateManyLaminaRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateLaminaService", "sincronizarFundos")
	return nil
}
