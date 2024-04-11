package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateExtratoService(extratos []domain.ExtratoDomain) *resterrors.RestErr {
	logger.Info("Init CreateFundosService", "sincronizarFundos")

	if err := fs.repository.CreateManyExtratoRepository(extratos); err != nil {
		logger.Error("Error calling CreateManyExtratoRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateFundosService", "sincronizarFundos")
	return nil
}
