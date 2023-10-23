package services

import "api-fundos-investimentos/configuration/logger"

func (fs *fundosDomainService) DownloadArquivosCVMService(folder string) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")
	fs.externo.DownloadArquivosCVMPort(folder)
	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}
