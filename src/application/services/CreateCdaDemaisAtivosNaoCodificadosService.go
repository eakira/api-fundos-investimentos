package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaDemaisAtivosNaoCodificadosService(demaisAtivos []domain.CdaDemaisAtivosNaoCodificadosDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaDemaisAtivosNaoCodificadosService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaDemaisAtivosNaoCodificadosRepository(demaisAtivos); err != nil {
		logger.Error("Error calling CreateManyCdaDemaisAtivosNaoCodificadosRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaDemaisAtivosNaoCodificadosService", "sincronizarFundos")
	return nil
}
