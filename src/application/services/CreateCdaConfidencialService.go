package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaConfidencialService(cda []domain.CdaConfidencialDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaConfidencialService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaConfidencialRepository(cda); err != nil {
		logger.Error("Error calling CreateManyCdaConfidencialRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaConfidencialService", "sincronizarFundos")
	return nil
}
