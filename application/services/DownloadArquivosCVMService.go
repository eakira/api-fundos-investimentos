package services

import "api-fundos-investimentos/configuration/logger"

func (fs *fundosDomainService) DownloadArquivosCVMService() {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")
	fs.externo.DownloadArquivosCVMPort()
	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}
