package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaSwapService(domain []domain.CdaSwapDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaSwapService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaSwapRepository(domain); err != nil {
		logger.Error("Error calling CreateManyCdaSwapRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaSwapService", "sincronizarFundos")
	return nil
}
