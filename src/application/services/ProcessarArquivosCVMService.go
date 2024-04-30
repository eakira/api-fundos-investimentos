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
	chanError := make(chan *resterrors.RestErr)
	var wg sync.WaitGroup
	wg.Add(1)

	go fs.processaArquivo(arquivosDomain, cabecalhoChan, linhaChan, jsonChan, mensagemChan, chanError, &wg)
	err := salvarProcessamento(fs, arquivosDomain)
	if err != nil {
		return err
	}

	for err := range chanError {
		defer wg.Done()
		return err // Retornar o erro recebido
	}
	wg.Wait()

	logger.Info("Processamento de Arquivos CVM Concluído", "sincronizarFundos")
	return nil
}

func (fs *fundosDomainService) processaArquivo(
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
	jsonChan chan []byte,
	mensagemChan chan response.FundosQueueResponse,
	chanError chan *resterrors.RestErr,
	wg *sync.WaitGroup,
) {
	go processaCsv(arquivosDomain, cabecalhoChan, linhaChan, chanError)
	go processarLinhas(arquivosDomain, cabecalhoChan, linhaChan, jsonChan, chanError)

	if env.GetPersistenciaLocal() {
		// Envio para persistência local
		fs.enviaPersistencia(jsonChan, mensagemChan, chanError)
	} else {
		// Envio para Kafka
		proximoQueue(jsonChan, mensagemChan, chanError)
		fs.queue.ProduceLote(mensagemChan)
	}
	defer wg.Done()

}

func processarLinhas(
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
	jsonChan chan []byte,
	chanError chan *resterrors.RestErr,
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
	defer close(jsonChan)
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

func mapeandoCollection(
	arquivosDomain domain.ArquivosDomain,
	mapCollection map[string]string,
) (
	collection string,
) {
	for key, value := range mapCollection {
		if strings.Contains(arquivosDomain.Endereco, key) {
			collection = value
			break
		}
	}
	return
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
	chanError chan *resterrors.RestErr,
) {
	reader, arquivo, error := openArquivo(arquivosDomain.Endereco)
	if error != nil {
		chanError <- error
		return
	}
	defer arquivo.Close()

	cabecalho, err := reader.Read()

	if err != nil {
		chanError <- resterrors.NewInternalServerError("Erro ao ler cabeçalho do CSV")
		logger.Error("Erro ao ler cabeçalho do CSV: ", err, "sincronizarFundos")
		return
	}
	cabecalhoChan <- cabecalho

	for {
		linha, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			chanError <- resterrors.NewInternalServerError("Erro ao ler linha do CSV")
			logger.Error("Erro ao ler linha do CSV: ", err, "sincronizarFundos")
			return
		}
		linhaChan <- linha
	}

	defer close(linhaChan)
	defer close(cabecalhoChan)

}

func openArquivo(
	endereco string,
) (
	reader *csv.Reader,
	arquivo *os.File,
	erro *resterrors.RestErr,
) {
	arquivo, err := os.Open(endereco)

	if err != nil {
		logger.Error("Erro ao abrir arquivo CSV: ", err, "sincronizarFundos")
		erro = resterrors.NewInternalServerError("Erro ao abrir arquivo CSV:")
		return
	}

	decoder := charmap.ISO8859_1.NewDecoder()
	reader = csv.NewReader(decoder.Reader(arquivo))
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	return
}

func salvarProcessamento(
	fs *fundosDomainService,
	arquivosDomain domain.ArquivosDomain,
) *resterrors.RestErr {
	arquivosDomain.UpdatedAt = time.Now()
	arquivosDomain.Processado = true
	arquivosDomain.Status = constants.PROCESSANDO

	return fs.repository.UpdateArquivosRepository(arquivosDomain)
}

func proximoQueue(
	jsonChan chan []byte,
	mensagemChan chan response.FundosQueueResponse,
	chanError chan *resterrors.RestErr,
) {
	for data := range jsonChan {
		response := response.FundosQueueResponse{
			Topic: env.GetTopicPersistenciaDados(),
			Queue: "update-all",
			Data:  data,
		}
		mensagemChan <- response
	}
	defer close(mensagemChan)
	defer close(chanError)
}

func (fs *fundosDomainService) enviaPersistencia(
	jsonChan chan []byte,
	mensagemChan chan response.FundosQueueResponse,
	chanError chan *resterrors.RestErr,
) *resterrors.RestErr {
	for data := range jsonChan {
		err := fs.CreateMany(data)
		if err != nil {
			return err
		}
	}
	defer close(mensagemChan)
	defer close(chanError)
	return nil
}
