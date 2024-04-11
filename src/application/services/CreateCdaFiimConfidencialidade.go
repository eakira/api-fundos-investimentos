package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaFiimConfidencialidade(domain []domain.CdaFiimConfidencialDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaFiimConfidencialidade", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaFiimConfidencialRepository(domain); err != nil {
		logger.Error("Error calling CreateManyCdaConfidencialRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaFiimConfidencialidade", "sincronizarFundos")
	return nil
}
