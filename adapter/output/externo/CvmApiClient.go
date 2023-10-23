package externo

import (
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"os"

	"github.com/go-resty/resty/v2"
)

var (
	client  *resty.Client
	CVM_URL = "CVM_URL"
)

type fundosClient struct{}

func NewFundosClient() *fundosClient {

	client = resty.New().SetBaseURL(os.Getenv(CVM_URL))
	return &fundosClient{}
}

func (nc *fundosClient) DownloadArquivosCVMPort() *resterrors.RestErr {
	logger.Info("Init DownloadArquivosCVMPort", "sincronizar")

	logger.Info("Finish DownloadArquivosCVMPort", "sincronizar")
	return nil
}
