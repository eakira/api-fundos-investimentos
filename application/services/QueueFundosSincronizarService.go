package services

import (
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/core/dto"
)

func (fs *fundosDomainService) QueueFundosSincronizarService(dto dto.FundosQueueDto) {
	logger.Info("Init QueueFundosExternoService", "sincronizarFundos")

	fs.queue.Produce(dto)
	logger.Info("Finish QueueFundosExternoService", "sincronizarFundos")
}
