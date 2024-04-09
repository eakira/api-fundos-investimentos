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

	arquivos := fs.externo.DownloadArquivosCVMPort(arquivosDomain.Endereco)

	salvandoDownload(fs, arquivosDomain)
	salvandoArquivos(fs, arquivosDomain, arquivos)

	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}

func salvandoDownload(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain) {
	arquivosDomain.UpdateAt = time.Now()
	arquivosDomain.Download = true
	arquivosDomain.Status = constants.FINALIZADO
	fs.repository.UpdateArquivosRepository(arquivosDomain)
}

func salvandoArquivos(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain, arquivos []string) {
	for _, endereco := range arquivos {
		arquivosDomain.CreatedAt = time.Now()
		arquivosDomain.Endereco = endereco
		arquivosDomain.Download = true
		arquivosDomain.Status = constants.DOWNLOAD
		proximo, _ := fs.repository.CreateArquivosRepository(arquivosDomain)
		proximoQueueDownload(fs, proximo)
	}

}

func proximoQueueDownload(fs *fundosDomainService, arquivosDomain *domain.ArquivosDomain) {
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
