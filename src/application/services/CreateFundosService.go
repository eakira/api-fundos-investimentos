package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateFundosService(fundos []domain.FundosDomain) *resterrors.RestErr {
	logger.Info("Init CreateFundosService", "sincronizarFundos")

	if err := fs.repository.CreateManyFundosRepository(fundos); err != nil {
		logger.Error("Error calling CreateManyFundosRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateFundosService", "sincronizarFundos")
	return nil
}
