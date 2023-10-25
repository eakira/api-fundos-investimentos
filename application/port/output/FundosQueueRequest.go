package output

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
)

type FundosPort interface {
	CreateArquivosRepository(
		arquivosDomain domain.ArquivosDomain,
	) (*domain.ArquivosDomain, *resterrors.RestErr)

	UpdateArquivosRepository(
		arquivosDomain domain.ArquivosDomain,
	) *resterrors.RestErr
}

type FundosExternoPort interface {
	DownloadArquivosCVMPort(string) *resterrors.RestErr
}

type FundosQueuePort interface {
	Produce(response.FundosQueueResponse) *resterrors.RestErr
}
