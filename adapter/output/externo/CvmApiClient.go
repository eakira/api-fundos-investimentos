package externo

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/go-resty/resty/v2"
)

var (
	client  *resty.Client
	CVM_URL = "CVM_URL"
)

type fundosClient struct{}

func NewFundosClient() *fundosClient {
	return &fundosClient{}
}

func (fc *fundosClient) DownloadArquivosCVMPort(file string) *resterrors.RestErr {
	logger.Info("Init DownloadArquivosCVMPort", "sincronizar")

	files := []string{
		os.Getenv(CVM_URL),
		file,
	}
	file = strings.Join(files, "")

	//Buscando o arquivo
	resp, err := http.Get(file)
	if err != nil {
		logger.Error("Erro ao tentar baixar o arquivo", err, "sincronizar")
		return resterrors.NewInternalServerError("Erro ao tentar baixar o arquivo")
	}
	defer resp.Body.Close()

	// Criando a pasta do arquivo
	pasta := env.GetPathArquivosCvm()
	err = os.MkdirAll(pasta, os.ModePerm)
	if err != nil {
		logger.Error("Erro ao tentar criar a pasta do arquivo", err, "sincronizar")
		return resterrors.NewInternalServerError("Erro ao tentar criar a pasta do arquivo")
	}

	filepath := []string{
		pasta,
		path.Base(resp.Request.URL.String()),
	}
	file = strings.Join(filepath, "")

	// Criando a pasta do arquivo
	out, err := os.Create(file)
	if err != nil {
		logger.Error("Erro ao criar o arquivo", err, "sincronizar")
		return resterrors.NewInternalServerError("Erro ao criar o arquivo")
	}
	defer out.Close()

	// Copiando o arquivo
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logger.Error("Erro ao salvar o arquivo", err, "sincronizar")
		return resterrors.NewInternalServerError("Erro ao salvar o arquivo")
	}
	logger.Info("Finish DownloadArquivosCVMPort", "sincronizar")

	return nil

}
