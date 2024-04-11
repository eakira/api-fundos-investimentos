package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaFiimService(fiim []domain.CdaFiimDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaFiimService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaFiimRepository(fiim); err != nil {
		logger.Error("Error calling CreateManyCdaConfidencialRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaFiimService", "sincronizarFundos")
	return nil
}
