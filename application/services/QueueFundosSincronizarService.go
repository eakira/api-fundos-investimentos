package services

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"fmt"
	"time"
)

func (fs *fundosDomainService) QueueFundosSincronizarService(tipo string) {
	logger.Info("Init QueueFundosExternoService", "sincronizarFundos")
	tipos := []string{
		"cadastros",
		"FIDC",
		"FIE",
		"FII",
		"FIP",
	}
	datas := getArquivosSufixo("2012-02-01", "2023-09-01", 0, 1, 0, "200601")

	logger.Info(
		fmt.Sprintf("json %v", datas),
		"sincronizar")
	for _, value := range tipos {
		response := response.FundosQueueResponse{
			Topic: env.GetTopicSincronizar(),
			Queue: "update-all",
			Data:  []string{value},
		}
		fs.queue.Produce(response)

	}

	logger.Info("Finish QueueFundosExternoService", "sincronizarFundos")
}

func getArquivosCadastros() []string {
	return []string{
		"FI/CAD/DADOS/cad_fi.csv",
		"FI/CAD/DADOS/cad_fi_hist.zip",
	}
}
func getArquivosBalancete() []string {
	return []string{
		"FI/CAD/DADOS/cad_fi.csv",
		"FI/CAD/DADOS/cad_fi_hist.zip",
	}
}

func getArquivosSufixo(dataInicio string, dataFinal string, ano int, mes int, dia int, formato string) []string {
	datas := []string{}
	dataInicioConvertida, _ := time.Parse("2006-01-02", dataInicio)
	dataFinalConvertida, _ := time.Parse("2006-01-02", dataFinal)
	for dataInicioConvertida.Before(dataFinalConvertida) {
		dataInicioConvertida = dataInicioConvertida.AddDate(ano, mes, dia)
		datas = append(datas, dataInicioConvertida.Format(formato))
	}

	return datas

}
