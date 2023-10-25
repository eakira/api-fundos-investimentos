package services

import (
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"time"
)

func (fs *fundosDomainService) DownloadArquivosCVMService(arquivosDomain domain.ArquivosDomain) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")
	fs.externo.DownloadArquivosCVMPort(arquivosDomain.Endereco)

	arquivosDomain.UpdateAt = time.Now()
	arquivosDomain.Download = true
	arquivosDomain.Status = constants.DOWNLOAD
	fs.repository.UpdateArquivosRepository(arquivosDomain)

	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}
