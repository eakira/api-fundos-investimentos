package input

import (
	"api-fundos-investimentos/adapter/output/model/response"
)

type FundosDomainService interface {
	QueueFundosSincronizarService(response.FundosQueueResponse)
	GetFundosExternoService()
	DownloadArquivosCVMService(folder string)
}
