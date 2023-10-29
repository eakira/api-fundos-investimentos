package services

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/encoding/charmap"
)

func (fs *fundosDomainService) ProcessarArquivosCVMService(arquivosDomain domain.ArquivosDomain) {
	logger.Info("Init GetFundosExternoService", "sincronizarFundos")

	processaArquivo(fs, arquivosDomain)

	salvandoProcessar(fs, arquivosDomain)
	logger.Info("Finish GetFundosExternoService", "sincronizarFundos")
}

func processaArquivo(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain) {

	cabecalhoChan := make(chan []string, 1)
	linhaChan := make(chan []string, 1000)
	jsonChan := make(chan []byte, 1000)
	mensagemChan := make(chan any, 1000)

	var wg sync.WaitGroup
	wg.Add(1)
	go processaCsv(linhaChan, cabecalhoChan, arquivosDomain)
	go processarLinha(linhaChan, cabecalhoChan, arquivosDomain, jsonChan)
	go converteLinha(jsonChan, arquivosDomain, mensagemChan)

	go fs.queue.ProduceLote(mensagemChan, &wg)

	wg.Wait()

}

func processarLinha(
	linhaChan chan []string,
	cabecalhoChan chan []string,
	arquivosDomain domain.ArquivosDomain,
	jsonChan chan []byte,
) {
	cabecalho := <-cabecalhoChan
	mapa := make(map[string]any)
	mapa["collection"] = arquivosDomain.TipoArquivo
	mapa["tipo-acao"] = "store"

	for linha := range linhaChan {
		for key, coluna := range cabecalho {
			mapa[coluna] = linha[key]
		}

		json, err := json.Marshal(mapa)
		if err != nil {
			fmt.Println("tre")
		}
		jsonChan <- json
	}
	close(jsonChan)

}

func processaCsv(
	linhaChan chan []string,
	cabecalhoChan chan []string,
	arquivosDomain domain.ArquivosDomain,
) {
	aquivo := strings.Replace(arquivosDomain.Endereco, ".zip", ".csv", 1)

	file, err := os.Open(env.GetPathArquivosCvm() + aquivo)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// reader := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(file))

	reader := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(file))
	//reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	cabecalhoChanArq, err := reader.Read()
	cabecalhoChan <- cabecalhoChanArq
	i := 0

	//	records, err := reader.ReadAll()
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.Error("error try reader.Read(): ", err, "processarArquivos")
			log.Fatal(err)
		}
		linhaChan <- record
		i++
		fmt.Println(i)
	}
	close(linhaChan)
	close(cabecalhoChan)
	fmt.Println("Done")
}

func salvandoProcessar(fs *fundosDomainService, arquivosDomain domain.ArquivosDomain) {
	arquivosDomain.UpdateAt = time.Now()
	arquivosDomain.Processado = true
	arquivosDomain.Status = constants.PROCESSANDO
	fs.repository.UpdateArquivosRepository(arquivosDomain)
}

func converteLinha(
	jsonChan chan []byte,
	arquivosDomain domain.ArquivosDomain,
	mensagemChan chan interface{},
) {

	for data := range jsonChan {
		switch arquivosDomain.TipoArquivo {
		case "cadastros":
			dados := response.FundosCadastrosResponse{}
			json.Unmarshal(data, &dados)
			mensagemChan <- dados

		case "balancete":
			dados := response.BalanceteResponse{}
			json.Unmarshal(data, &dados)
			mensagemChan <- dados

		case "cda":
			//		São vários arquivos precisa verificar quais arquivos vou usar

		case "informacoes-complementares":
			//		São vários arquivos precisa verificar quais arquivos vou usar

		case "extrato":
			dados := response.ExtratoResponse{}
			json.Unmarshal(data, &dados)
			mensagemChan <- dados

		case "informacao-diaria":
			dados := response.InformacaoDiariaResponse{}
			json.Unmarshal(data, &dados)
			mensagemChan <- dados

		case "lamina":
			//		São vários arquivos precisa verificar quais arquivos vou usar

		case "perfil-mensal":
			//		files = getFilesName(env.GetConfigCvmArquivosPerfilMensal())

		}

	}
	close(mensagemChan)

}
