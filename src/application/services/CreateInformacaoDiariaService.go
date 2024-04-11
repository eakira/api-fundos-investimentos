package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateInformacaoDiariaService(informacaoDiaria []domain.InformacaoDiariaDomain) *resterrors.RestErr {
	logger.Info("Init CreateFundosService", "sincronizarFundos")

	if err := fs.repository.CreateManyInformacaoDiariaRepository(informacaoDiaria); err != nil {
		logger.Error("Error calling CreateManyInformacaoDiariaRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateFundosService", "sincronizarFundos")
	return nil
}
