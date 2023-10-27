package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"

	"github.com/jinzhu/copier"
)

func (fc *fundosControllerInterface) CreateFundosController(request request.FundosCvmArquivosQueueRequest) {
	logger.Info("Init CreateFundosController", "sincronizarFundos")

	arquivoDomain := &domain.ArquivosDomain{}
	copier.Copy(arquivoDomain, request)
	fc.service.ProcessarArquivosCVMService(*arquivoDomain)

}
