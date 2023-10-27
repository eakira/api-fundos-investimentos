package input

import (
	"api-fundos-investimentos/application/domain"
)

type FundosDomainService interface {
	QueueFundosSincronizarService(string)
	GetFundosExternoService()
	DownloadArquivosCVMService(domain.ArquivosDomain)
	ProcessarArquivosCVMService(domain.ArquivosDomain)
	CreateFundosService(domain.FundosDomain)
	CreateBalanceteService(domain.BalanceteDomain)
}
