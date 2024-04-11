package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaFundosInvestimentosService(fundosInvestimentos []domain.CdaFundosInvestimentosDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaFundosInvestimentosService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaFundosInvestimentosRepository(fundosInvestimentos); err != nil {
		logger.Error("Error calling CreateManyCdaFundosInvestimentosRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaFundosInvestimentosService", "sincronizarFundos")
	return nil
}
