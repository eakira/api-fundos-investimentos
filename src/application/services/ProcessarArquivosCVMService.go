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
	logger.Info("Iniciando Processamento de Arquivos CVM", "ProcessarArquivosCVMService")

	cabecalhoChan := make(chan []string, 1)
	linhaChan := make(chan []string, 100000)
	jsonChan := make(chan []byte, 5000)
	mensagemChan := make(chan response.FundosQueueResponse, 1000)

	var wg sync.WaitGroup
	defer close(cabecalhoChan)
	defer close(linhaChan)
	defer close(jsonChan)
	defer close(mensagemChan)

	wg.Add(1)
	go processaArquivo(fs, arquivosDomain, cabecalhoChan, linhaChan, jsonChan, mensagemChan, &wg)

	if err := salvarProcessamento(fs, arquivosDomain); err != nil {
		logger.Error("Erro ao salvar processamento: ", err, "ProcessarArquivosCVMService")
		return resterrors.NewInternalServerError("Erro ao salvar processamento")
	}

	wg.Wait()
	logger.Info("Processamento de Arquivos CVM Concluído", "ProcessarArquivosCVMService")
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
	defer wg.Done()

	if env.GetPersistenciaLocal() {
		// Envio para persistência local
		if err := processaLocal(fs, arquivosDomain, cabecalhoChan, linhaChan, jsonChan); err != nil {
			logger.Error("Erro ao processar localmente: ", err, "processaArquivo")
		}
	} else {
		// Envio para Kafka
		processaKafka(fs, jsonChan, mensagemChan)
	}
}

func processaLocal(
	fs *fundosDomainService,
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
	jsonChan chan []byte,
) *resterrors.RestErr {
	defer close(jsonChan)

	processaCsv(arquivosDomain, cabecalhoChan, linhaChan)

	// Processar linhas e enviar JSON
	tipoArquivo := definirCollection(arquivosDomain)
	limit := env.GetLimitInsert()
	mapaJson := make([]map[string]interface{}, 0, limit)

	for linha := range linhaChan {
		mapa := make(map[string]interface{})
		mapa["collection"] = tipoArquivo
		mapa["tipo-acao"] = "store"

		for key, coluna := range <-cabecalhoChan {
			if key < len(linha) {
				mapa[coluna] = linha[key]
			}
		}

		mapaJson = append(mapaJson, mapa)

		if len(mapaJson) == limit {
			if err := enviarJSON(mapaJson, jsonChan); err != nil {
				logger.Error("Erro ao serializar JSON: ", err, "processaLocal")
			}
			mapaJson = make([]map[string]interface{}, 0, limit)
		}
	}

	// Enviar o restante dos dados, se houver
	if len(mapaJson) > 0 {
		if err := enviarJSON(mapaJson, jsonChan); err != nil {
			logger.Error("Erro ao serializar JSON: ", err, "processaLocal")
		}
	}

	return nil
}

func processaKafka(
	fs *fundosDomainService,
	jsonChan chan []byte,
	mensagemChan chan response.FundosQueueResponse,
) {
	defer close(mensagemChan)

	for data := range jsonChan {
		response := response.FundosQueueResponse{
			Topic: env.GetTopicPersistenciaDados(),
			Queue: "update-all",
			Data:  data,
		}
		mensagemChan <- response
	}
}

func processaCsv(
	arquivosDomain domain.ArquivosDomain,
	cabecalhoChan chan []string,
	linhaChan chan []string,
) {
	nomeArquivo := strings.Replace(arquivosDomain.Endereco, ".zip", ".csv", 1)
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		logger.Error("Erro ao abrir arquivo CSV: ", err, "ProcessarCsv")
		return
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
		return
	}
	cabecalhoChan <- cabecalho

	for {
		linha, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("Erro ao ler linha do CSV: ", err, "ProcessarCsv")
			break
		}
		linhaChan <- linha
	}

	close(linhaChan)
}

func salvarProcessamento(
	fs *fundosDomainService,
	arquivosDomain domain.ArquivosDomain,
) *resterrors.RestErr {
	arquivosDomain.UpdateAt = time.Now()
	arquivosDomain.Processado = true
	arquivosDomain.Status = constants.PROCESSANDO

	if err := fs.repository.UpdateArquivosRepository(arquivosDomain); err != nil {
		logger.Error("Erro ao salvar processamento: ", err, "SalvarProcessamento")
		return resterrors.NewInternalServerError("Erro ao salvar processamento")
	}

	return nil
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
