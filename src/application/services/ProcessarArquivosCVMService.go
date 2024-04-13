package services

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
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
func (fs *fundosDomainService) ProcessarArquivosCVMService(arquivosDomain domain.ArquivosDomain) *resterrors.RestErr {
	logger.Info("Iniciando Processamento de Arquivos CVM", "sincronizarFundos")

	cabecalhoChan := make(chan []string, 1)
	linhaChan := make(chan []string, 100000)
	jsonChan := make(chan []byte, 5000)
	mensagemChan := make(chan response.FundosQueueResponse, 1000)

	var wg sync.WaitGroup
	wg.Add(1)

	go processaArquivo(fs, arquivosDomain, cabecalhoChan, linhaChan, jsonChan, mensagemChan, &wg)
	salvarProcessamento(fs, arquivosDomain)
	wg.Wait()

	logger.Info("Processamento de Arquivos CVM Concluído", "sincronizarFundos")
	return nil
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

	if env.GetPersistenciaLocal() {
		// Envio para persistência local
		enviaPersistencia(fs, jsonChan, mensagemChan, wg)
	} else {
		// Envio para Kafka
		proximoQueue(jsonChan, mensagemChan)
		fs.queue.ProduceLote(mensagemChan, wg)
	}

}

func processarLinhas(
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
	jsonChan chan []byte,
) {
	cabecalho := <-cabecalhoChan
	tipoArquivo := definirCollection(arquivosDomain)
	limit := env.GetLimitInsert()
	mapaJson := make([]map[string]interface{}, 0, limit)

	for linha := range linhaChan {
		mapa := make(map[string]interface{})
		mapa["collection"] = tipoArquivo
		mapa["tipo-acao"] = "store"

		for key, coluna := range cabecalho {
			if key < len(linha) {
				mapa[coluna] = linha[key]
			}
		}

		mapaJson = append(mapaJson, mapa)

		if len(mapaJson) == limit {
			if err := enviarJSON(mapaJson, jsonChan); err != nil {
				logger.Error("Erro ao serializar JSON: ", err, "sincronizarFundos")
			}
			mapaJson = make([]map[string]interface{}, 0, limit)
		}
	}

	// Enviar o restante dos dados, se houver
	if len(mapaJson) > 0 {
		if err := enviarJSON(mapaJson, jsonChan); err != nil {
			logger.Error("Erro ao serializar JSON: ", err, "sincronizarFundos")
		}
	}

	close(jsonChan)
}

func definirCollection(
	arquivosDomain domain.ArquivosDomain,
) string {
	tipoArquivo := arquivosDomain.TipoArquivo

	switch tipoArquivo {
	case "cda":
		mapCollection := env.GetMapCda()
		return mapeandoCollection(arquivosDomain, mapCollection)

	case "informacoes-complementares":
		mapCollection := env.GetMapInformacaoComplementar()
		return mapeandoCollection(arquivosDomain, mapCollection)

	case "lamina":
		mapCollection := env.GetMapLamina()
		return mapeandoCollection(arquivosDomain, mapCollection)

	default:
		return tipoArquivo
	}
}

func mapeandoCollection(arquivosDomain domain.ArquivosDomain, mapCollection map[string]string) string {
	collection := ""
	for key, value := range mapCollection {
		if strings.Contains(arquivosDomain.Endereco, key) {
			collection = value
			break
		}
	}
	return collection
}

func enviarJSON(mapaJson []map[string]interface{}, jsonChan chan []byte) error {
	jsonData, err := json.Marshal(mapaJson)
	if err != nil {
		return err
	}
	jsonChan <- jsonData
	return nil
}

func processaCsv(
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
) {
	nomeArquivo := strings.Replace(arquivosDomain.Endereco, ".zip", ".csv", 1)
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		logger.Error("Erro ao abrir arquivo CSV: ", err, "sincronizarFundos")
	}
	defer arquivo.Close()

	decoder := charmap.ISO8859_1.NewDecoder()
	reader := csv.NewReader(decoder.Reader(arquivo))
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	cabecalho, err := reader.Read()
	if err != nil {
		logger.Error("Erro ao ler cabeçalho do CSV: ", err, "sincronizarFundos")
	}
	cabecalhoChan <- cabecalho

	for {
		linha, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("Erro ao ler linha do CSV: ", err, "sincronizarFundos")
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
		logger.Error("Erro ao salvar processamento: ", err, "sincronizarFundos")
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

func enviaPersistencia(
	fs *fundosDomainService,
	jsonChan chan []byte,
	mensagemChan chan response.FundosQueueResponse,
	wg *sync.WaitGroup,
) {
	for data := range jsonChan {
		CreateMany(fs, data)
	}
	close(mensagemChan)
	wg.Done()
}
