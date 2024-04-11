package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) CreateInformacaoComplementarFundoService(domain []domain.InformacoesFundoDomain) {
	logger.Info("Init CreateInformacaoComplementarFundo", "sincronizarFundos")

	fs.repository.CreateManyInformacaoComplementarFundoRepository(domain)
	logger.Info("Finish CreateInformacaoComplementarFundo", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateInformacaoComplementarDivulgacaoService(domain []domain.InformacoesDivulgacaoDomain) {
	logger.Info("Init CreateInformacaoComplementarDivulgacao", "sincronizarFundos")

	fs.repository.CreateManyInformacaoComplementarDivulgacaoRepository(domain)
	logger.Info("Finish CreateInformacaoComplementarDivulgacao", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateInformacaoComplementarCotistaService(domain []domain.InformacoesCotistaDomain) {
	logger.Info("Init CreateInformacaoComplementarCotista", "sincronizarFundos")

	fs.repository.CreateManyInformacaoComplementarCotistaRepository(domain)
	logger.Info("Finish CreateInformacaoComplementarCotista", "sincronizarFundos")
}

func (fs *fundosDomainService) CreateInformacaoComplementarServicoPrestadoService(domain []domain.ServicoPrestadoDomain) {
	logger.Info("Init CreateInformacaoComplementarServicoPrestado", "sincronizarFundos")

	fs.repository.CreateManyInformacaoComplementarServicoPrestadoRepository(domain)
	logger.Info("Finish CreateInformacaoComplementarServicoPrestado", "sincronizarFundos")
}
