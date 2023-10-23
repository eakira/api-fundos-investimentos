package controller

import (
	"api-fundos-investimentos/configuration/logger"
)

func (fc *fundosControllerInterface) DownloadArquivosCVMController() {
	logger.Info("Init DownloadArquivosCVMController", "sincronizarFundos")
	fc.service.DownloadArquivosCVMService()
	logger.Info("Finish DownloadArquivosCVMController", "sincronizarFundos")
}
