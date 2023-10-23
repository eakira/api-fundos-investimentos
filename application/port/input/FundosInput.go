package input

import "api-fundos-investimentos/adapter/output/model/request"

type FundosDomainService interface {
	QueueFundosSincronizarService(request.FundosQueueRequest)
	GetFundosExternoService()
	DownloadArquivosCVMService()
}
