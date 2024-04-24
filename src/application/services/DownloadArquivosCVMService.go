package services

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
	"time"

	"github.com/jinzhu/copier"
)

func (fs *fundosDomainService) DownloadArquivosCVMService(arquivosDomain domain.ArquivosDomain) *resterrors.RestErr {
	logger.Info("Init DownloadArquivosCVMService", "sincronizarFundos")

	arquivos, err := fs.externo.DownloadArquivosCVMPort(arquivosDomain.Endereco, arquivosDomain.Baixar)
	if err != nil {
		return err
	}

	err = salvandoDownload(fs, arquivosDomain)
	if err != nil {
		return err
	}

	err = salvandoArquivos(fs, arquivosDomain, arquivos)
	if err != nil {
		return err
	}

	logger.Info("Finish DownloadArquivosCVMService", "sincronizarFundos")
	return nil
}

func salvandoDownload(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain) *resterrors.RestErr {
	arquivosDomain.UpdatedAt = time.Now()
	arquivosDomain.Download = true
	arquivosDomain.Status = constants.FINALIZADO
	return fs.repository.UpdateArquivosRepository(arquivosDomain)
}

func salvandoArquivos(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain, arquivos []string) *resterrors.RestErr {
	for _, endereco := range arquivos {
		arquivosDomain.CreatedAt = time.Now()
		arquivosDomain.Endereco = endereco
		arquivosDomain.Download = true
		arquivosDomain.Status = constants.DOWNLOAD
		proximo, err := fs.repository.CreateArquivosRepository(arquivosDomain)
		if err != nil {
			return err
		}

		err = proximoQueueDownload(fs, proximo)
		if err != nil {
			return err
		}
	}
	return nil
}

func proximoQueueDownload(fs *fundosDomainService, arquivosDomain *domain.ArquivosDomain) *resterrors.RestErr {
	arquivosRequest := &request.FundosCvmArquivosQueueRequest{}
	copier.Copy(arquivosRequest, arquivosDomain)

	data, _ := json.Marshal(arquivosRequest)
	response := response.FundosQueueResponse{
		Topic: env.GetTopicProcessarArquivos(),
		Queue: "update-all",
		Data:  data,
	}

	return fs.queue.Produce(response)
}
