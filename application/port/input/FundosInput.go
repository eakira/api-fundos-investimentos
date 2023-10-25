package input

import "api-fundos-investimentos/adapter/output/model/request"

type FundosDomainService interface {
	QueueFundosSincronizarService(string)
	GetFundosExternoService()
	DownloadArquivosCVMService(request.FundosDownloadCvmFilesQueueRequest)
}
