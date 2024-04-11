package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaDemaisAtivosService(domain []domain.CdaDemaisAtivosDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaDemaisAtivosService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaDemaisAtivosRepository(domain); err != nil {
		logger.Error("Error calling CreateManyCdaSwapRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaDemaisAtivosService", "sincronizarFundos")
	return nil
}
