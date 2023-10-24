package controller

import (
	"api-fundos-investimentos/configuration/logger"
)

func (fc *fundosControllerInterface) DownloadArquivosCVMController(file string) {
	logger.Info("Init DownloadArquivosCVMController", "sincronizarFundos")
	fc.service.DownloadArquivosCVMService(file)
	logger.Info("Finish DownloadArquivosCVMController", "sincronizarFundos")
}
