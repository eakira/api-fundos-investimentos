package services

import (
	"api-fundos-investimentos/configuration/logger"
	"fmt"
	"strings"
	"time"
)

func (fs *fundosDomainService) QueueFundosSincronizarService(tipo string) {
	logger.Info("Init QueueFundosExternoService", "sincronizarFundos")
	files := []string{}
	logger.Info(time.Now().AddDate(0, -1, 0).Format("2006-01-02"), "sincronizarFundos")

	switch tipo {
	case "cadastros":
		files = getArquivosCadastros()
	case "balancete":
		files = getArquivosBalancete()
	case "cda":
		files = getArquivosCda()
	case "informacoes-complementares":
		files = getArquivosInformacoesComplementares()
	case "extrato":
		files = getArquivosExtrato()
	case "informacao-diaria":
		files = getArquivosInformacaoDiaria()
	case "lamina":
		files = getArquivosLamina()
	case "perfil-mensal":
		files = getArquivosPerfilMensal()
	}

	logger.Info(
		fmt.Sprintf("json %v", files),
		"sincronizar")
	/*
		for _, value := range files {
			response := response.FundosQueueResponse{
				Topic: env.GetTopicSincronizar(),
				Queue: "update-all",
				Data:  []string{value},
			}
			fs.queue.Produce(response)

		}
	*/
	logger.Info("Finish QueueFundosExternoService", "sincronizarFundos")
}

func getArquivosCadastros() []string {
	return []string{
		"FI/CAD/DADOS/cad_fi.csv",
		"FI/CAD/DADOS/cad_fi_hist.zip",
	}
}

func getArquivosPerfilMensal() []string {
	logger.Info("Init getArquivosLamina", "sincronizarFundos")

	files := getFilesName(
		"FI/DOC/PERFIL_MENSAL/DADOS/perfil_mensal_fi_",
		".zip",
		"2019-01-01",
		time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
		"200601",
		0,
		1,
		0,
	)

	logger.Info("Finish getArquivosLamina", "sincronizarFundos")
	return files
}

func getArquivosLamina() []string {
	logger.Info("Init getArquivosLamina", "sincronizarFundos")

	files := getFilesName(
		"FI/DOC/LAMINA/DADOS/lamina_fi_",
		".zip",
		"2019-01-01",
		time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
		"200601",
		0,
		1,
		0,
	)

	histfiles := getFilesName(
		"FI/DOC/LAMINA/DADOS/HIST/lamina_fi_",
		".zip",
		"2014-01-01",
		"2018-12-01",
		"200601",
		0,
		1,
		0,
	)

	logger.Info("Finish getArquivosLamina", "sincronizarFundos")
	return append(files, histfiles...)
}

func getArquivosInformacaoDiaria() []string {
	logger.Info("Init getArquivosInformacaoDiaria", "sincronizarFundos")

	files := getFilesName(
		"FI/DOC/INF_DIARIO/DADOS/inf_diario_fi_",
		".zip",
		"2021-01-01",
		time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
		"200601",
		0,
		1,
		0,
	)

	histfiles := getFilesName(
		"FI/DOC/INF_DIARIO/DADOS/HIST/inf_diario_fi_",
		".zip",
		"2000-01-01",
		"2019-12-01",
		"2006",
		1,
		0,
		0,
	)

	logger.Info("Finish getArquivosInformacaoDiaria", "sincronizarFundos")
	return append(files, histfiles...)
}

func getArquivosInformacoesComplementares() []string {
	logger.Info("Init getArquivosInformacoesComplementares", "sincronizarFundos")

	files := getFilesName(
		"FI/DOC/COMPL/DADOS/compl_fi_",
		".zip",
		"2018-01-01",
		"2019-04-01",
		"200601",
		0,
		1,
		0,
	)

	logger.Info("Finish getArquivosInformacoesComplementares", "sincronizarFundos")
	return files
}
func getArquivosExtrato() []string {
	logger.Info("Init getArquivosExtrato", "sincronizarFundos")
	files := []string{"FI/DOC/EXTRATO/DADOS/extrato_fi.csv"}
	histfiles := getFilesName(
		"FI/DOC/EXTRATO/DADOS/extrato_fi_",
		".csv",
		"2015-01-01",
		"2022-12-31",
		"2006",
		1,
		0,
		0,
	)

	logger.Info("Finish getArquivosExtrato", "sincronizarFundos")
	return append(files, histfiles...)
}

func getArquivosCda() []string {
	logger.Info("Init getArquivosCda", "sincronizarFundos")

	files := getFilesName(
		"FI/DOC/CDA/DADOS/cda_fi_",
		".zip",
		"2021-01-01",
		time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
		"200601",
		0,
		1,
		0,
	)

	histfiles := getFilesName(
		"FI/DOC/CDA/DADOS/HIST/cda_fi_",
		".zip",
		"2005-01-01",
		"2020-12-01",
		"2006",
		1,
		0,
		0,
	)

	logger.Info("Finish getArquivosCda", "sincronizarFundos")
	return append(files, histfiles...)
}

func getArquivosBalancete() []string {
	logger.Info("Init getArquivosBalancete", "sincronizarFundos")

	files := getFilesName(
		"FI/DOC/BALANCETE/DADOS/balancete_fi_",
		".zip",
		"2015-01-01",
		time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
		"200601",
		0,
		1,
		0,
	)

	histfiles := getFilesName(
		"FI/DOC/BALANCETE/DADOS/HIST/balancete_fi_",
		".zip",
		"2005-05-01",
		"2014-12-01",
		"200601",
		0,
		1,
		0,
	)

	logger.Info("Finish getArquivosBalancete", "sincronizarFundos")
	return append(files, histfiles...)
}

func getFilesName(folder, extension, DataInicial, DataFinal, formato string, ano, mes, dia int) []string {
	logger.Info("Init getFilesName", "sincronizarFundos")
	sufixos := getArquivosSufixo(DataInicial, DataFinal, ano, mes, dia, formato)

	files := []string{}
	for _, value := range sufixos {
		array := []string{
			folder,
			value,
			extension,
		}
		files = append(files, strings.Join(array, ""))
	}

	logger.Info("Finish getFilesName", "sincronizarFundos")

	return files

}

func getArquivosSufixo(dataInicio string, dataFinal string, ano int, mes int, dia int, formato string) []string {
	datas := []string{}
	dataInicioConvertida, _ := time.Parse("2006-01-02", dataInicio)
	dataFinalConvertida, _ := time.Parse("2006-01-02", dataFinal)
	for dataInicioConvertida.Before(dataFinalConvertida) {
		datas = append(datas, dataInicioConvertida.Format(formato))
		dataInicioConvertida = dataInicioConvertida.AddDate(ano, mes, dia)
	}
	datas = append(datas, dataInicioConvertida.Format(formato))

	return datas

}
