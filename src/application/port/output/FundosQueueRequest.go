package output

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"sync"
)

type FundosPort interface {
	CreateArquivosRepository(
		arquivosDomain domain.ArquivosDomain,
	) (*domain.ArquivosDomain, *resterrors.RestErr)

	UpdateArquivosRepository(
		arquivosDomain domain.ArquivosDomain,
	) *resterrors.RestErr

	CreateFundosRepository(
		fundosDomain domain.FundosDomain,
	) (*domain.FundosDomain, *resterrors.RestErr)

	CreateManyFundosRepository(
		fundosDomain []domain.FundosDomain,
	) *resterrors.RestErr

	CreateManyBalecenteRepository(
		balanceteDomain []domain.BalanceteDomain,
	) *resterrors.RestErr

	CreateManyExtratoRepository(
		extratoDomain []domain.ExtratoDomain,
	) *resterrors.RestErr

	CreateManyInformacaoDiariaRepository(
		informacaoDiariaDomain []domain.InformacaoDiariaDomain,
	) *resterrors.RestErr
}

type FundosExternoPort interface {
	DownloadArquivosCVMPort(string) []string
}

type FundosQueuePort interface {
	Produce(response.FundosQueueResponse) *resterrors.RestErr
	ProduceLote(chan response.FundosQueueResponse, *sync.WaitGroup) *resterrors.RestErr
}
