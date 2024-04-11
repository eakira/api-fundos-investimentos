package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateBalanceteService(balancetes []domain.BalanceteDomain) *resterrors.RestErr {
	logger.Info("Init CreateFundosService", "sincronizarFundos")

	if err := fs.repository.CreateManyBalecenteRepository(balancetes); err != nil {
		logger.Error("Error calling CreateManyBalecenteRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateFundosService", "sincronizarFundos")
	return nil
}
