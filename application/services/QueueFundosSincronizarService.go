package services

import "api-fundos-investimentos/configuration/logger"

func (fs *fundosDomainService) QueueFundosSincronizarService(topic string, message string) {
	logger.Info("Init QueueFundosExternoService", "sincronizarFundos")
	fs.queue.Produce(topic, message)
	logger.Info("Finish QueueFundosExternoService", "sincronizarFundos")
}
