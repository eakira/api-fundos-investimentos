package services

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"encoding/json"
	"time"

	"github.com/jinzhu/copier"
)

func (fs *fundosDomainService) DownloadArquivosCVMService(arquivosDomain domain.ArquivosDomain) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")

	err := fs.externo.DownloadArquivosCVMPort(arquivosDomain.Endereco)
	if err != nil {
		logger.Error("Error trying to DownloadArquivosCVMPort", err, "sincronizarFundos")
		return
	}

	salvandoDownload(fs, arquivosDomain)

	proximoQueueDownload(fs, arquivosDomain)

	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}

func salvandoDownload(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain) {
	arquivosDomain.UpdateAt = time.Now()
	arquivosDomain.Download = true
	arquivosDomain.Status = constants.DOWNLOAD
	fs.repository.UpdateArquivosRepository(arquivosDomain)
}

func proximoQueueDownload(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain) {
	arquivosRequest := &request.FundosCvmArquivosQueueRequest{}
	copier.Copy(arquivosRequest, arquivosDomain)

	data, _ := json.Marshal(arquivosRequest)
	response := response.FundosQueueResponse{
		Topic: env.GetTopicProcessarArquivos(),
		Queue: "update-all",
		Data:  data,
	}
	fs.queue.Produce(response)

}
