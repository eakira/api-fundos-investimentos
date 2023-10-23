package output

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/configuration/resterrors"
)

type FundosPort interface {
}

type FundosExternoPort interface {
	DownloadArquivosCVMPort(string) *resterrors.RestErr
}

type FundosQueuePort interface {
	Produce(response.FundosQueueResponse) *resterrors.RestErr
}
