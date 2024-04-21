package input

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
)

type FundosDomainService interface {
	QueueFundosSincronizarService(string, bool) *resterrors.RestErr
	DownloadArquivosCVMService(domain.ArquivosDomain) *resterrors.RestErr
	ProcessarArquivosCVMService(domain.ArquivosDomain) *resterrors.RestErr
	CreateFundosService([]domain.FundosDomain) *resterrors.RestErr
	CreateBalanceteService([]domain.BalanceteDomain) *resterrors.RestErr
	CreateExtratoService([]domain.ExtratoDomain) *resterrors.RestErr
	CreateInformacaoDiariaService([]domain.InformacaoDiariaDomain) *resterrors.RestErr
	CreateCdaConfidencialService([]domain.CdaConfidencialDomain) *resterrors.RestErr
	CreatePerfilMensalService([]domain.PerfilMensalDomain) *resterrors.RestErr
	CreateCdaSelicService([]domain.CdaSelicDomain) *resterrors.RestErr
	CreateCdaFundosInvestimentosService([]domain.CdaFundosInvestimentosDomain) *resterrors.RestErr
	CreateCdaSwapService([]domain.CdaSwapDomain) *resterrors.RestErr
	CreateCdaDemaisAtivosService([]domain.CdaDemaisAtivosDomain) *resterrors.RestErr
	CreateCdaDepositoAPrazoOutrosAtivosService([]domain.CdaDepositoAPrazoOutrosAtivosDomain) *resterrors.RestErr
	CreateCdaAgroCreditoPrivadoService([]domain.CdaAgroCreditoPrivadoDomain) *resterrors.RestErr
	CreateCdaInvestimentosExteriorService([]domain.CdaInvestimentosExteriorDomain) *resterrors.RestErr
	CreateCdaDemaisAtivosNaoCodificadosService([]domain.CdaDemaisAtivosNaoCodificadosDomain) *resterrors.RestErr
	CreateCdaFiimConfidencialidade([]domain.CdaFiimConfidencialDomain) *resterrors.RestErr
	CreateCdaPatrominioLiquido([]domain.CdaPatrimonioLiquidoDomain) *resterrors.RestErr

	CreateCdaFiimService([]domain.CdaFiimDomain) *resterrors.RestErr
	CreateInformacaoComplementarFundoService([]domain.InformacoesFundoDomain) *resterrors.RestErr
	CreateInformacaoComplementarDivulgacaoService([]domain.InformacoesDivulgacaoDomain) *resterrors.RestErr
	CreateInformacaoComplementarCotistaService([]domain.InformacoesCotistaDomain) *resterrors.RestErr
	CreateInformacaoComplementarServicoPrestadoService([]domain.ServicoPrestadoDomain) *resterrors.RestErr
	CreateLaminaService([]domain.LaminaDomain) *resterrors.RestErr
	CreateLaminaCarteiraService([]domain.LaminaCarteiraDomain) *resterrors.RestErr
	CreateLaminaRentabilidadeAnoService([]domain.LaminaRentabilidadeAnoDomain) *resterrors.RestErr
	CreateLaminaRentabilidadeMesService([]domain.LaminaRentabilidadeMesDomain) *resterrors.RestErr
}
