package services

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jinzhu/copier"
)

func (fs *fundosDomainService) ProcessarArquivosCVMService(arquivosDomain domain.ArquivosDomain) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")

	switch arquivosDomain.TipoArquivo {
	case "cadastros":
		processaArquivoCadastro(arquivosDomain)
	case "balancete":

	case "cda":

	case "informacoes-complementares":

	case "extrato":

	case "informacao-diaria":

	case "lamina":

	case "perfil-mensal":

	}

	salvandoProcessar(fs, arquivosDomain)

	//proximoQueue(fs, arquivosDomain)

	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}
func processaArquivoCadastro(arquivosDomain domain.ArquivosDomain) {
	file, err := os.Open(env.GetPathArquivosCvm() + arquivosDomain.Endereco)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
		i++
	}
	fmt.Println(i)
}

func salvandoProcessar(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain) {
	arquivosDomain.UpdateAt = time.Now()
	arquivosDomain.Processado = true
	arquivosDomain.Status = constants.PROCESSANDO
	fs.repository.UpdateArquivosRepository(arquivosDomain)
}

func proximoQueue(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain) {
	arquivosRequest := request.FundosCvmArquivosQueueRequest{}
	copier.Copy(arquivosRequest, arquivosDomain)

	data, _ := json.Marshal(arquivosRequest)
	response := response.FundosQueueResponse{
		Topic: env.GetTopicSincronizar(),
		Queue: "update-all",
		Data:  data,
	}
	fs.queue.Produce(response)

}
