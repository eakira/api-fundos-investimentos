package input

import "api-fundos-investimentos/core/dto"

type FundosDomainService interface {
	QueueFundosSincronizarService(dto.FundosQueueDto)
	GetFundosExternoService()
	DownloadArquivosCVMService()
}
