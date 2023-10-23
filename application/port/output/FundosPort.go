package output

import (
	"api-fundos-investimentos/configuration/resterrors"
	"api-fundos-investimentos/core/dto"
)

type FundosPort interface {
}

type FundosExternoPort interface {
	DownloadArquivosCVMPort() *resterrors.RestErr
}

type FundosQueuePort interface {
	Produce(dto.FundosQueueDto) *resterrors.RestErr
}
