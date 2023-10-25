package services

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) DownloadArquivosCVMService(request request.FundosDownloadCvmFilesQueueRequest) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")

	//	fs.externo.DownloadArquivosCVMPort(request)
	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}
