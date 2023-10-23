package services

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) QueueFundosSincronizarService(response response.FundosQueueResponse) {
	logger.Info("Init QueueFundosExternoService", "sincronizarFundos")

	fs.queue.Produce(response)
	logger.Info("Finish QueueFundosExternoService", "sincronizarFundos")
}
