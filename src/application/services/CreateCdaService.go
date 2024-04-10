package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreateCdaBlc1Service(domain []domain.CdaBlc1Domain) {
	logger.Info("Init CreateCdaBlc1Service", "sincronizarFundos")

	fs.repository.CreateManyCdaBlc1Repository(domain)
	logger.Info("Finish CreateCdaBlc1Service", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaBlc2Service(domain []domain.CdaBlc2Domain) {
	logger.Info("Init CreateCdaBlc2Service", "sincronizarFundos")

	fs.repository.CreateManyCdaBlc2Repository(domain)
	logger.Info("Finish CreateCdaBlc2Service", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaBlc3Service(domain []domain.CdaBlc3Domain) {
	logger.Info("Init CreateCdaBlc3Service", "sincronizarFundos")

	fs.repository.CreateManyCdaBlc3Repository(domain)
	logger.Info("Finish CreateCdaBlc3Service", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaBlc4Service(domain []domain.CdaBlc4Domain) {
	logger.Info("Init CreateCdaBlc4Service", "sincronizarFundos")

	fs.repository.CreateManyCdaBlc4Repository(domain)
	logger.Info("Finish CreateCdaBlc4Service", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaBlc5Service(domain []domain.CdaBlc5Domain) {
	logger.Info("Init CreateCdaBlc5Service", "sincronizarFundos")

	fs.repository.CreateManyCdaBlc5Repository(domain)
	logger.Info("Finish CreateCdaBlc5Service", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaBlc6Service(domain []domain.CdaBlc6Domain) {
	logger.Info("Init CreateCdaBlc6Service", "sincronizarFundos")

	fs.repository.CreateManyCdaBlc6Repository(domain)
	logger.Info("Finish CreateCdaBlc6Service", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaBlc7Service(domain []domain.CdaBlc7Domain) {
	logger.Info("Init CreateCdaBlc7Service", "sincronizarFundos")

	fs.repository.CreateManyCdaBlc7Repository(domain)
	logger.Info("Finish CreateCdaBlc7Service", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaBlc8Service(domain []domain.CdaBlc8Domain) {
	logger.Info("Init CreateCdaBlc8Service", "sincronizarFundos")

	fs.repository.CreateManyCdaBlc8Repository(domain)
	logger.Info("Finish CreateCdaBlc8Service", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaConfidencialService(domain []domain.CdaConfidencialDomain) {
	logger.Info("Init CreateCdaConfidencialService", "sincronizarFundos")

	fs.repository.CreateManyCdaConfidencialRepository(domain)
	logger.Info("Finish CreateCdaConfidencialService", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaFiimService(domain []domain.CdaFiimDomain) {
	logger.Info("Init CreateCdaFiimService", "sincronizarFundos")

	fs.repository.CreateManyCdaFiimRepository(domain)
	logger.Info("Finish CreateCdaFiimService", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaFiimConfidencialidade(domain []domain.CdaFiimConfidencialDomain) {
	logger.Info("Init CreateCdaFiimConfidencialidade", "sincronizarFundos")

	fs.repository.CreateManyCdaFiimConfidencialRepository(domain)
	logger.Info("Finish CreateCdaFiimConfidencialidade", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateCdaPatrominioLiquido(domain []domain.CdaPatrimonioLiquidoDomain) {
	logger.Info("Init CreateCdaPatrominioLiquido", "sincronizarFundos")

	fs.repository.CreateManyCdaPatrimonioLiquidoRepository(domain)
	logger.Info("Finish CreateCdaPatrominioLiquido", "sincronizarFundos")
}
