package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaDepositoAPrazoOutrosAtivosService(domain []domain.CdaDepositoAPrazoOutrosAtivosDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaDepositoAPrazoOutrosAtivosService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaDepositoAPrazoOutrosAtivosRepository(domain); err != nil {
		logger.Error("Error calling CreateManyCdaSwapRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaDepositoAPrazoOutrosAtivosService", "sincronizarFundos")
	return nil
}
