package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
)

func (fc *fundosControllerInterface) DownloadArquivosCVMController(request request.FundosDownloadCvmFilesQueueRequest) {
	logger.Info("Init DownloadArquivosCVMController", "sincronizarFundos")
	arquivoDomain := domain.ArquivosDomain{
		Endereco:    request.FileName,
		TipoArquivo: request.Tipo,
		Referencia:  request.Referencia,
	}

	fc.service.DownloadArquivosCVMService(arquivoDomain)

}
