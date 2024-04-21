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

	CreateManyPerfilMensalRepository(
		[]domain.PerfilMensalDomain,
	) *resterrors.RestErr

	CreateManyCdaSelicRepository(
		[]domain.CdaSelicDomain,
	) *resterrors.RestErr

	CreateManyCdaFundosInvestimentosRepository(
		[]domain.CdaFundosInvestimentosDomain,
	) *resterrors.RestErr

	CreateManyCdaSwapRepository(
		[]domain.CdaSwapDomain,
	) *resterrors.RestErr

	CreateManyCdaDemaisAtivosRepository(
		[]domain.CdaDemaisAtivosDomain,
	) *resterrors.RestErr

	CreateManyCdaDepositoAPrazoOutrosAtivosRepository(
		[]domain.CdaDepositoAPrazoOutrosAtivosDomain,
	) *resterrors.RestErr

	CreateManyCdaAgroCreditoPrivadoRepository(
		[]domain.CdaAgroCreditoPrivadoDomain,
	) *resterrors.RestErr

	CreateManyCdaInvestimentosExteriorRepository(
		[]domain.CdaInvestimentosExteriorDomain,
	) *resterrors.RestErr

	CreateManyCdaDemaisAtivosNaoCodificadosRepository(
		[]domain.CdaDemaisAtivosNaoCodificadosDomain,
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

	CreateManyLaminaRepository(
		[]domain.LaminaDomain,
	) *resterrors.RestErr

	CreateManyLaminaCarteiraRepository(
		[]domain.LaminaCarteiraDomain,
	) *resterrors.RestErr

	CreateManyLaminaRentabilidadeAnoRepository(
		[]domain.LaminaRentabilidadeAnoDomain,
	) *resterrors.RestErr

	CreateManyLaminaRentabilidadeMesRepository(
		[]domain.LaminaRentabilidadeMesDomain,
	) *resterrors.RestErr
}

type FundosExternoPort interface {
	DownloadArquivosCVMPort(string, bool) []string
}

type FundosQueuePort interface {
	Produce(response.FundosQueueResponse) *resterrors.RestErr
	ProduceLote(chan response.FundosQueueResponse, *sync.WaitGroup) *resterrors.RestErr
}
