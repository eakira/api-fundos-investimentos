package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaInvestimentosExteriorService(investimentosExterior []domain.CdaInvestimentosExteriorDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaInvestimentosExteriorService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaInvestimentosExteriorRepository(investimentosExterior); err != nil {
		logger.Error("Error calling CreateManyCdaInvestimentosExteriorRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaInvestimentosExteriorService", "sincronizarFundos")
	return nil
}
