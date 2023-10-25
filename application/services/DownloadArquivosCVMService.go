package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fs *fundosDomainService) DownloadArquivosCVMService(arquivosDomain domain.ArquivosDomain) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")

	//	fs.externo.DownloadArquivosCVMPort(request)
	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}
