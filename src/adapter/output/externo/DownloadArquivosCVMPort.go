package externo

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func (fc *fundosClient) DownloadArquivosCVMPort(file string) []string {
	logger.Info("Iniciando DownloadArquivosCVMPort", "sincronizar")
	var nomes []string

	storagePath := env.GetPathArquivosCvm() + path.Dir(file)

	downloadArquivo(storagePath, file)

	// Descompactando o arquivo ZIP, se necessário
	if path.Ext(file) == ".zip" {
		nomes = unzip(storagePath + "/" + path.Base(file))
	}

	return nomes
}

func downloadArquivo(storagePath, file string) {
	url := env.GetCvmUrl() + file

	// Fazendo a requisição HTTP para baixar o arquivo
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Erro ao tentar baixar o arquivo:", err, "sincronizar")
	}
	defer resp.Body.Close()

	// Criando o diretório de armazenamento do arquivo, se não existir
	err = os.MkdirAll(storagePath, os.ModePerm)
	if err != nil {
		logger.Error("Erro ao tentar criar a pasta do arquivo:", err, "sincronizar")
	}

	// Criando o arquivo local
	out, err := os.Create(storagePath + "/" + path.Base(file))
	if err != nil {
		logger.Error("Erro ao criar o arquivo:", err, "sincronizar")
	}
	defer out.Close()

	// Copiando o conteúdo do arquivo baixado para o arquivo local
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logger.Error("Erro ao salvar o arquivo:", err, "sincronizar")
	}
	logger.Info("Download de arquivo concluído com sucesso.", "sincronizar")
}

// unzip descompacta um arquivo ZIP especificado e salva seus conteúdos no sistema de arquivos local.
func unzip(filename string) []string {
	// Abre o arquivo ZIP para leitura
	reader, err := zip.OpenReader(filename)
	if err != nil {
		logger.Error(fmt.Sprintf("Erro ao abrir arquivo ZIP: %s", filename), err, "sincronizar")
	}
	defer reader.Close()

	var nomes []string

	// Itera sobre todos os arquivos no arquivo ZIP
	for _, file := range reader.File {
		// Abre o arquivo dentro do arquivo ZIP
		in, err := file.Open()
		if err != nil {
			logger.Error("Erro ao abrir arquivo dentro do ZIP:", err, "sincronizar")
			continue
		}
		defer in.Close()

		// Cria o arquivo correspondente no sistema de arquivos local
		relname := path.Join(path.Dir(filename), file.Name)
		dir := path.Dir(relname)
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			logger.Error("Erro ao criar diretório para o arquivo ZIP:", err, "sincronizar")
			continue
		}

		out, err := os.Create(relname)
		if err != nil {
			logger.Error("Erro ao criar arquivo dentro do ZIP:", err, "sincronizar")
			continue
		}
		defer out.Close()

		// Copia o conteúdo do arquivo dentro do arquivo ZIP para o arquivo local
		_, err = io.Copy(out, in)
		if err != nil {
			logger.Error("Erro ao copiar conteúdo do arquivo dentro do ZIP:", err, "sincronizar")
			continue
		}

		nomes = append(nomes, file.Name)
	}

	return nomes
}
