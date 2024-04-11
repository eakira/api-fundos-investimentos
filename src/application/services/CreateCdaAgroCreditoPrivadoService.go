package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaAgroCreditoPrivadoService(agro []domain.CdaAgroCreditoPrivadoDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaAgroCreditoPrivadoService", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaAgroCreditoPrivadoRepository(agro); err != nil {
		logger.Error("Error calling CreateManyCdaSwapRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaAgroCreditoPrivadoService", "sincronizarFundos")
	return nil
}
