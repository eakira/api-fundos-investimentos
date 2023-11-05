package externo

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path"
)

func (fc *fundosClient) DownloadArquivosCVMPort(file string) *resterrors.RestErr {
	logger.Info("Init DownloadArquivosCVMPort", "sincronizar")

	url := os.Getenv(CVM_URL) + file

	//Buscando o arquivo
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Erro ao tentar baixar o arquivo", err, "sincronizar")
		return resterrors.NewInternalServerError("Erro ao tentar baixar o arquivo")
	}
	defer resp.Body.Close()

	// Criando a pasta do arquivo
	storage := env.GetPathArquivosCvm()
	storage = storage + path.Dir(file)
	err = os.MkdirAll(storage, os.ModePerm)
	if err != nil {
		logger.Error("Erro ao tentar criar a pasta do arquivo", err, "sincronizar")
		return resterrors.NewInternalServerError("Erro ao tentar criar a pasta do arquivo")
	}

	// Criando a pasta do arquivo
	out, err := os.Create(storage + "/" + path.Base(file))
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

	if file[len(file)-3:] == "zip" {
		unzip(storage + "/" + path.Base(file))
	}

	return nil

}

func unzip(filename string) {
	reader, _ := zip.OpenReader(filename)
	defer reader.Close()
	for _, file := range reader.File {
		in, _ := file.Open()
		defer in.Close()
		relname := path.Join(path.Dir(filename), file.Name)
		dir := path.Dir(relname)
		os.MkdirAll(dir, 0777)
		out, _ := os.Create(relname)
		defer out.Close()
		io.Copy(out, in)
	}
}
