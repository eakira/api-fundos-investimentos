package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreateInformaComplementarFundoService(domain []domain.InformacoesFundoDomain) {
	logger.Info("Init CreateInformaComplementarFundo", "sincronizarFundos")

	fs.repository.CreateInformaComplementarFundoRepository(domain)
	logger.Info("Finish CreateInformaComplementarFundo", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateInformaComplementarDivulgacaoService(domain []domain.InformacoesDivulgacaoDomain) {
	logger.Info("Init CreateInformaComplementarDivulgacao", "sincronizarFundos")

	fs.repository.CreateInformaComplementarDivulgacaoRepository(domain)
	logger.Info("Finish CreateInformaComplementarDivulgacao", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateInformaComplementarCotistaService(domain []domain.InformacoesCotistaDomain) {
	logger.Info("Init CreateInformaComplementarCotista", "sincronizarFundos")

	fs.repository.CreateInformaComplementarCotistaRepository(domain)
	logger.Info("Finish CreateInformaComplementarCotista", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateInformaComplementarServicoPrestadoService(domain []domain.ServicoPrestadoDomain) {
	logger.Info("Init CreateInformaComplementarServicoPrestado", "sincronizarFundos")

	fs.repository.CreateInformaComplementarServicoPrestadoRepository(domain)
	logger.Info("Finish CreateInformaComplementarServicoPrestado", "sincronizarFundos")
}
