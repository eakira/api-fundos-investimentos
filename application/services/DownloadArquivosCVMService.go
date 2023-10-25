package services

import (
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"time"
)

func (fs *fundosDomainService) DownloadArquivosCVMService(arquivosDomain domain.ArquivosDomain) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")
	err := fs.externo.DownloadArquivosCVMPort(arquivosDomain.Endereco)
	if err != nil {
		logger.Error("Error trying to DownloadArquivosCVMPort", err, "sincronizarFundos")
		return
	}

	arquivosDomain.UpdateAt = time.Now()
	arquivosDomain.Download = true
	arquivosDomain.Status = constants.DOWNLOAD
	fs.repository.UpdateArquivosRepository(arquivosDomain)

	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}
