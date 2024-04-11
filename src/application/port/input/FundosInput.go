package input

import (
	"api-fundos-investimentos/application/domain"
)

type FundosDomainService interface {
	QueueFundosSincronizarService(string)
	GetFundosExternoService()
	DownloadArquivosCVMService(domain.ArquivosDomain)
	ProcessarArquivosCVMService(domain.ArquivosDomain)
	CreateFundosService([]domain.FundosDomain)
	CreateBalanceteService([]domain.BalanceteDomain)
	CreateExtratoService([]domain.ExtratoDomain)
	CreateInformacaoDiariaService([]domain.InformacaoDiariaDomain)
	CreateCdaBlc1Service([]domain.CdaBlc1Domain)
	CreateCdaBlc2Service([]domain.CdaBlc2Domain)
	CreateCdaBlc3Service([]domain.CdaBlc3Domain)
	CreateCdaBlc4Service([]domain.CdaBlc4Domain)
	CreateCdaBlc5Service([]domain.CdaBlc5Domain)
	CreateCdaBlc6Service([]domain.CdaBlc6Domain)
	CreateCdaBlc7Service([]domain.CdaBlc7Domain)
	CreateCdaBlc8Service([]domain.CdaBlc8Domain)
	CreateInformaComplementarFundoService([]domain.InformacoesFundoDomain)
	CreateInformaComplementarDivulgacaoService([]domain.InformacoesDivulgacaoDomain)
	CreateInformaComplementarCotistaService([]domain.InformacoesCotistaDomain)
	CreateInformaComplementarServicoPrestadoService([]domain.ServicoPrestadoDomain)
}
