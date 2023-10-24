package services

import "api-fundos-investimentos/configuration/logger"

func (fs *fundosDomainService) DownloadArquivosCVMService(file string) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")

	fs.externo.DownloadArquivosCVMPort(file)
	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}
