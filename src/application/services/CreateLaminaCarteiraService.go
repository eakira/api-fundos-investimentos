package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateLaminaCarteiraService(laminaCarteira []domain.LaminaCarteiraDomain) *resterrors.RestErr {
	logger.Info("Init CreateLaminaCarteiraService", "sincronizarFundos")

	if err := fs.repository.CreateManyLaminaCarteiraRepository(laminaCarteira); err != nil {
		logger.Error("Error calling CreateManyLaminaCarteiraRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateLaminaCarteiraService", "sincronizarFundos")
	return nil
}
