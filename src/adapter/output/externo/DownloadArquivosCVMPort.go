package externo

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	permissions = 0777
)

func (fc *fundosClient) DownloadArquivosCVMPort(file string, baixar bool) ([]string, *resterrors.RestErr) {
	logger.Info("Iniciando DownloadArquivosCVMPort", "sincronizarFundos")

	storagePath := filepath.Join(env.GetPathArquivosCvm(), filepath.Dir(file))
	err := os.MkdirAll(storagePath, permissions)
	if err != nil {
		logger.Error("Erro ao tentar criar a pasta do arquivo:", err, "sincronizarFundos")
		return nil, resterrors.NewInternalServerError("Erro ao tentar criar a pasta do arquivo")
	}

	localFilePath := filepath.Join(storagePath, filepath.Base(file))
	if baixar {
		err = downloadArquivo(env.GetCvmUrl()+file, localFilePath)
		if err != nil {
			logger.Error("Erro ao baixar o arquivo:", err, "sincronizarFundos")
			return nil, resterrors.NewInternalServerError("Erro ao baixar o arquivo")
		}
	}

	var nomes []string
	if filepath.Ext(file) == ".zip" {
		nomes = unzip(localFilePath)
	} else {
		nomes = append(nomes, localFilePath)
	}

	logger.Info("Download de arquivo concluído com sucesso.", "sincronizarFundos")
	return nomes, nil
}

func downloadArquivo(
	url,
	localFilePath string,
) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erro ao tentar baixar o arquivo: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao salvar o arquivo: %v", err)
	}

	return nil
}

func unzip(filename string) []string {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		logger.Error(fmt.Sprintf("Erro ao abrir arquivo ZIP: %s", filename), err, "sincronizarFundos")
		return nil
	}
	defer reader.Close()

	var nomes []string

	for _, file := range reader.File {
		in, err := file.Open()
		if err != nil {
			logger.Error("Erro ao abrir arquivo dentro do ZIP:", err, "sincronizarFundos")
			continue
		}
		defer in.Close()

		relname := filepath.Join(filepath.Dir(filename), file.Name)
		err = os.MkdirAll(filepath.Dir(relname), permissions)
		if err != nil {
			logger.Error("Erro ao criar diretório para o arquivo ZIP:", err, "sincronizarFundos")
			continue
		}

		out, err := os.Create(relname)
		if err != nil {
			logger.Error("Erro ao criar arquivo dentro do ZIP:", err, "sincronizarFundos")
			continue
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			logger.Error("Erro ao copiar conteúdo do arquivo dentro do ZIP:", err, "sincronizarFundos")
			continue
		}

		nomes = append(nomes, relname)
	}

	return nomes
}
