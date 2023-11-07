package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) ProcessarArquivosCVMController(request request.FundosCvmArquivosQueueRequest) {
	logger.Info("Init ProcessarArquivosCVMController", "sincronizarFundos")

	arquivoDomain := &domain.ArquivosDomain{}
	copier.Copy(arquivoDomain, request)
	fc.service.ProcessarArquivosCVMService(*arquivoDomain)

}