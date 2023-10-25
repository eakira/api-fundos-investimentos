package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
)

func (fc *fundosControllerInterface) DownloadArquivosCVMController(request request.FundosDownloadCvmFilesQueueRequest) {
	logger.Info("Init DownloadArquivosCVMController", "sincronizarFundos")
	fc.service.DownloadArquivosCVMService(request)
	logger.Info("Finish DownloadArquivosCVMController", "sincronizarFundos")
}
