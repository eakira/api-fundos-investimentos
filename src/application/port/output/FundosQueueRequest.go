package output

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"sync"
)

type FundosPort interface {
	CreateArquivosRepository(
		domain.ArquivosDomain,
	) (*domain.ArquivosDomain, *resterrors.RestErr)

	UpdateArquivosRepository(
		domain.ArquivosDomain,
	) *resterrors.RestErr

	CreateFundosRepository(
		domain.FundosDomain,
	) (*domain.FundosDomain, *resterrors.RestErr)

	CreateManyFundosRepository(
		[]domain.FundosDomain,
	) *resterrors.RestErr

	CreateManyBalecenteRepository(
		[]domain.BalanceteDomain,
	) *resterrors.RestErr

	CreateManyExtratoRepository(
		[]domain.ExtratoDomain,
	) *resterrors.RestErr

	CreateManyInformacaoDiariaRepository(
		[]domain.InformacaoDiariaDomain,
	) *resterrors.RestErr

	CreateManyCdaBlc1Repository(
		[]domain.CdaBlc1Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc2Repository(
		[]domain.CdaBlc2Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc3Repository(
		[]domain.CdaBlc3Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc4Repository(
		[]domain.CdaBlc4Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc5Repository(
		[]domain.CdaBlc5Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc6Repository(
		[]domain.CdaBlc6Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc7Repository(
		[]domain.CdaBlc7Domain,
	) *resterrors.RestErr

	CreateManyCdaBlc8Repository(
		[]domain.CdaBlc8Domain,
	) *resterrors.RestErr

	CreateManyCdaConfidencialRepository(
		[]domain.CdaConfidencialDomain,
	) *resterrors.RestErr

	CreateManyCdaFiimRepository(
		[]domain.CdaFiimDomain,
	) *resterrors.RestErr

	CreateManyCdaFiimConfidencialRepository(
		[]domain.CdaFiimConfidencialDomain,
	) *resterrors.RestErr

	CreateManyCdaPatrimonioLiquidoRepository(
		[]domain.CdaPatrimonioLiquidoDomain,
	) *resterrors.RestErr

	CreateManyInformacaoComplementarFundoRepository(
		[]domain.InformacoesFundoDomain,
	) *resterrors.RestErr

	CreateManyInformacaoComplementarDivulgacaoRepository(
		[]domain.InformacoesDivulgacaoDomain,
	) *resterrors.RestErr

	CreateManyInformacaoComplementarCotistaRepository(
		[]domain.InformacoesCotistaDomain,
	) *resterrors.RestErr

	CreateManyInformacaoComplementarServicoPrestadoRepository(
		[]domain.ServicoPrestadoDomain,
	) *resterrors.RestErr
}

type FundosExternoPort interface {
	DownloadArquivosCVMPort(string) []string
}

type FundosQueuePort interface {
	Produce(response.FundosQueueResponse) *resterrors.RestErr
	ProduceLote(chan response.FundosQueueResponse, *sync.WaitGroup) *resterrors.RestErr
}
