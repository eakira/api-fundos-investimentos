package controller

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/core/dto"
)

func (fc *fundosControllerInterface) DownloadArquivosCvm() {
	logger.Info("Init DownloadArquivosCvm", "sincronizarFundos")
	dto := dto.FundosQueueDto{
		Topic: env.GetTopicSincronizar(),
		Queue: "update-all",
		Data:  nil,
	}

	fc.service.QueueFundosSincronizarService(dto)
	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}
