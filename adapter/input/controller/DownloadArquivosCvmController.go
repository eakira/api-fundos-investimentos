package controller

import (
	"api-fundos-investimentos/configuration/logger"
)

func (fc *fundosControllerInterface) DownloadArquivosCVMController(folder string) {
	logger.Info("Init DownloadArquivosCVMController", "sincronizarFundos")
	fc.service.DownloadArquivosCVMService(folder)
	logger.Info("Finish DownloadArquivosCVMController", "sincronizarFundos")
}
