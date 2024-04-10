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

	CreateManyCdaBlc1Repository(
		cdaDomain []domain.CdaBlc1Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc2Repository(
		cdaDomain []domain.CdaBlc2Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc3Repository(
		cdaDomain []domain.CdaBlc3Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc4Repository(
		cdaDomain []domain.CdaBlc4Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc5Repository(
		cdaDomain []domain.CdaBlc5Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc6Repository(
		cdaDomain []domain.CdaBlc6Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc7Repository(
		cdaDomain []domain.CdaBlc7Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc8Repository(
		cdaDomain []domain.CdaBlc8Domain,
	) *resterrors.RestErr

	CreateManyCdaConfidencialRepository(
		cdaDomain []domain.CdaConfidencialDomain,
	) *resterrors.RestErr

	CreateManyCdaFiimRepository(
		cdaDomain []domain.CdaFiimDomain,
	) *resterrors.RestErr

	CreateManyCdaFiimConfidencialRepository(
		cdaDomain []domain.CdaFiimConfidencialDomain,
	) *resterrors.RestErr

	CreateManyCdaPatrimonioLiquidoRepository(
		cdaDomain []domain.CdaPatrimonioLiquidoDomain,
	) *resterrors.RestErr
}

type FundosExternoPort interface {
	DownloadArquivosCVMPort(string) []string
}

type FundosQueuePort interface {
	Produce(response.FundosQueueResponse) *resterrors.RestErr
	ProduceLote(chan response.FundosQueueResponse, *sync.WaitGroup) *resterrors.RestErr
}
