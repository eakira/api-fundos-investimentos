package services

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"encoding/json"
	"strings"
	"time"
)

func (fs *fundosDomainService) QueueFundosSincronizarService(tipo string) {
	logger.Info("Init QueueFundosExternoService", "sincronizarFundos")
	files := []response.FundosDownloadCvmFilesQueueResponse{}

	switch tipo {
	case "cadastros":
		files = getArquivosCadastro(env.GetConfigCvmArquivosCadastros())

	case "balancete":
		files = getFilesName(env.GetConfigCvmBalancete())
		histfiles := getFilesName(env.GetConfigCvmBalanceteHist())
		files = append(files, histfiles...)

	case "cda":
		files = getFilesName(env.GetConfigCvmCda())
		histfiles := getFilesName(env.GetConfigCvmCdaHist())
		files = append(files, histfiles...)

	case "informacoes-complementares":
		files = getFilesName(env.GetConfigCvmInformacoesComplementares())

	case "extrato":
		files = getFilesName(env.GetConfigCvmExtrato())

	case "informacao-diaria":
		files = getFilesName(env.GetConfigInformacaoDiaria())
		histfiles := getFilesName(env.GetConfigInformacaoDiariaHist())
		files = append(files, histfiles...)

	case "lamina":
		files = getFilesName(env.GetConfigCvmArquivosLaminas())
		histfiles := getFilesName(env.GetConfigCvmArquivosLaminasHist())
		files = append(files, histfiles...)

	case "perfil-mensal":
		files = getFilesName(env.GetConfigCvmArquivosPerfilMensal())

	}

	for _, value := range files {
		value.CreatedAt = time.Now()
		value.UpdateAt = time.Now()
		value.Status = constants.ENVIADO

		data, _ := json.Marshal(value)
		response := response.FundosQueueResponse{
			Topic: env.GetTopicSincronizar(),
			Queue: "update-all",
			Data:  data,
		}
		fs.queue.Produce(response)

	}

	logger.Info("Finish QueueFundosExternoService", "sincronizarFundos")
}

func getFilesName(arquivos env.ArquivosCVM) []response.FundosDownloadCvmFilesQueueResponse {
	logger.Info("Init getFilesName", "sincronizarFundos")
	sufixos := getArquivosSufixo(arquivos)
	files := []response.FundosDownloadCvmFilesQueueResponse{}
	for _, value := range sufixos {
		array := []string{
			arquivos.Folder,
			value,
			arquivos.Extension,
		}
		file := response.FundosDownloadCvmFilesQueueResponse{
			Endereco:    strings.Join(array, ""),
			Referencia:  value,
			TipoArquivo: arquivos.Tipo,
		}
		files = append(files, file)
	}

	logger.Info("Finish getFilesName", "sincronizarFundos")

	return files

}

func getArquivosSufixo(arquivo env.ArquivosCVM) []string {

	sufixos := []string{}
	dataInicioConvertida, _ := time.Parse("2006-01-02", arquivo.DataInicial)
	dataFinalConvertida, _ := time.Parse("2006-01-02", arquivo.DataFinal)
	for dataInicioConvertida.Before(dataFinalConvertida) {
		sufixos = append(sufixos, dataInicioConvertida.Format(arquivo.Formato))
		dataInicioConvertida = dataInicioConvertida.AddDate(arquivo.Ano, arquivo.Mes, arquivo.Dia)
	}
	return append(sufixos, dataInicioConvertida.Format(arquivo.Formato))

}

func getArquivosCadastro(env []string) []response.FundosDownloadCvmFilesQueueResponse {
	files := []response.FundosDownloadCvmFilesQueueResponse{}

	file := response.FundosDownloadCvmFilesQueueResponse{
		Endereco:    env[0],
		Referencia:  "2023",
		TipoArquivo: env[2],
	}
	files = append(files, file)
	file = response.FundosDownloadCvmFilesQueueResponse{
		Endereco:    env[0],
		Referencia:  "2022",
		TipoArquivo: env[2],
	}
	return append(files, file)
}
