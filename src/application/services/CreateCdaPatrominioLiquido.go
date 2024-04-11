package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateCdaPatrominioLiquido(domain []domain.CdaPatrimonioLiquidoDomain) *resterrors.RestErr {
	logger.Info("Init CreateCdaPatrominioLiquido", "sincronizarFundos")

	if err := fs.repository.CreateManyCdaPatrimonioLiquidoRepository(domain); err != nil {
		logger.Error("Error calling CreateManyCdaConfidencialRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateCdaPatrominioLiquido", "sincronizarFundos")
	return nil
}
