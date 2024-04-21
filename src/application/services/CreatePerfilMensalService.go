package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreatePerfilMensalService(perfilMensalDomain []domain.PerfilMensalDomain) *resterrors.RestErr {
	logger.Info("Init CreatePerfilMensalService", "sincronizarFundos")

	if err := fs.repository.CreateManyPerfilMensalRepository(perfilMensalDomain); err != nil {
		logger.Error("Error calling CreateManyPerfilMensalRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreatePerfilMensalService", "sincronizarFundos")
	return nil
}
