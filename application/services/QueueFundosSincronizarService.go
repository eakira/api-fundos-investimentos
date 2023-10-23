package services

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) QueueFundosSincronizarService(request request.FundosQueueRequest) {
	logger.Info("Init QueueFundosExternoService", "sincronizarFundos")

	fs.queue.Produce(request)
	logger.Info("Finish QueueFundosExternoService", "sincronizarFundos")
}
