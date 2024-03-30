package services

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/encoding/charmap"
)

// ProcessarArquivosCVMService processa arquivos CVM
func (fs *fundosDomainService) ProcessarArquivosCVMService(arquivosDomain domain.ArquivosDomain) {
	logger.Info("Iniciando Processamento de Arquivos CVM", "ProcessarArquivosCVMService")

	cabecalhoChan := make(chan []string, 1)
	linhaChan := make(chan []string, 100000)
	jsonChan := make(chan []byte, 500)
	mensagemChan := make(chan response.FundosQueueResponse, 1000)

	var wg sync.WaitGroup
	wg.Add(1)

	go processaArquivo(fs, arquivosDomain, cabecalhoChan, linhaChan, jsonChan, mensagemChan, &wg)

	wg.Wait()

	salvarProcessamento(fs, arquivosDomain)

	logger.Info("Processamento de Arquivos CVM Concluído", "ProcessarArquivosCVMService")
}

func processaArquivo(
	fs *fundosDomainService,
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
	jsonChan chan []byte,
	mensagemChan chan response.FundosQueueResponse,
	wg *sync.WaitGroup,
) {

	go processaCsv(arquivosDomain, cabecalhoChan, linhaChan)

	go processarLinhas(arquivosDomain, cabecalhoChan, linhaChan, jsonChan)

	go proximoQueue(jsonChan, mensagemChan)

	fs.queue.ProduceLote(mensagemChan, wg)
}

func processarLinhas(
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
	jsonChan chan []byte,
) {
	cabecalho := <-cabecalhoChan

	// Use um tamanho de buffer adequado para a slice
	limit := env.GetLimitInsert()
	if len(cabecalho) > 20 {
		limit = 1
	}

	mapaJson := make([]map[string]any, 0, limit)

	for linha := range linhaChan {
		mapa := make(map[string]any)
		mapa["collection"] = arquivosDomain.TipoArquivo
		mapa["tipo-acao"] = "store"

		for key, coluna := range cabecalho {
			if key < len(linha) {
				mapa[coluna] = linha[key]
			}
		}
		mapaJson = append(mapaJson, mapa)

		if len(mapaJson) == limit {
			json, err := json.Marshal(mapaJson)
			if err != nil {
				logger.Error("Erro ao serializar JSON: ", err, "ProcessarLinhas")
				continue
			}
			jsonChan <- json
			mapaJson = make([]map[string]any, 0, limit)
		}
	}

	if len(mapaJson) > 0 {
		json, err := json.Marshal(mapaJson)
		if err != nil {
			logger.Error("Erro ao serializar JSON: ", err, "ProcessarLinhas")
			panic("Erro ao serializar JSON")
		}
		jsonChan <- json
	}

	close(jsonChan)
}

func processaCsv(
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
) {
	nomeArquivo := strings.Replace(arquivosDomain.Endereco, ".zip", ".csv", 1)
	arquivo, err := os.Open(env.GetPathArquivosCvm() + nomeArquivo)

	if err != nil {
		logger.Error("Erro ao abrir arquivo CSV: ", err, "ProcessarCsv")
		panic(err)
	}
	defer arquivo.Close()

	decoder := charmap.ISO8859_1.NewDecoder()
	reader := csv.NewReader(decoder.Reader(arquivo))
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	cabecalho, err := reader.Read()
	if err != nil {
		logger.Error("Erro ao ler cabeçalho do CSV: ", err, "ProcessarCsv")
		panic(err)
	}
	cabecalhoChan <- cabecalho

	for {
		linha, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("Erro ao ler linha do CSV: ", err, "ProcessarCsv")
			panic(err)
		}
		linhaChan <- linha
	}

	close(linhaChan)
	close(cabecalhoChan)
}

func salvarProcessamento(
	fs *fundosDomainService,
	arquivosDomain domain.ArquivosDomain,
) {
	arquivosDomain.UpdateAt = time.Now()
	arquivosDomain.Processado = true
	arquivosDomain.Status = constants.PROCESSANDO

	err := fs.repository.UpdateArquivosRepository(arquivosDomain)
	if err != nil {
		logger.Error("Erro ao salvar processamento: ", err, "SalvarProcessamento")
		panic(err)
	}
}

func proximoQueue(
	jsonChan chan []byte,
	mensagemChan chan response.FundosQueueResponse,
) {
	for data := range jsonChan {
		response := response.FundosQueueResponse{
			Topic: env.GetTopicPersistenciaDados(),
			Queue: "update-all",
			Data:  data,
		}
		mensagemChan <- response
	}
	close(mensagemChan)
}
