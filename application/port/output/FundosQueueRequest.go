package output

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/resterrors"
)

type FundosPort interface {
}

type FundosExternoPort interface {
	DownloadArquivosCVMPort() *resterrors.RestErr
}

type FundosQueuePort interface {
	Produce(request.FundosQueueRequest) *resterrors.RestErr
}
