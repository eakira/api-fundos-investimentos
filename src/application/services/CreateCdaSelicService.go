package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaSelicService(selic []domain.CdaSelicDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaSelicService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaSelicRepository(selic); err != nil {
		logger.Error("Error calling CreateManyCdaSelicRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaSelicService", "sincronizarFundos")
	return nil
}
