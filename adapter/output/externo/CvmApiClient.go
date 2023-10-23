package externo

import (
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	cp "github.com/otiai10/copy"
)

var (
	client  *resty.Client
	CVM_URL = "CVM_URL"
)

type fundosClient struct{}

func NewFundosClient() *fundosClient {
	return &fundosClient{}
}

func (fc *fundosClient) DownloadArquivosCVMPort(folder string) *resterrors.RestErr {
	logger.Info("Init DownloadArquivosCVMPort", "sincronizar")

	folders := []string{
		os.Getenv(CVM_URL),
		folder,
	}
	folder = strings.Join(folders, "")
	logger.Info(folder, "sincronizar")

	err := cp.Copy("https://dados.cvm.gov.br/dados/FIP/DOC/INF_TRIMESTRAL/DADOS/", "storage/CVM")
	if err != nil {
		logger.Error("Erro ao copiar os arquivos", err, "sincronizar")
		return resterrors.NewInternalServerError("Erro ao copiar os arquivo")
	}

	logger.Info("Finish DownloadArquivosCVMPort", "sincronizar")
	return nil
}
