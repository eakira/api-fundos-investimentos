package input

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
)

type FundosDomainService interface {
	QueueFundosSincronizarService(string) *resterrors.RestErr
	GetFundosExternoService() *resterrors.RestErr
	DownloadArquivosCVMService(domain.ArquivosDomain) *resterrors.RestErr
	ProcessarArquivosCVMService(domain.ArquivosDomain) *resterrors.RestErr
	CreateFundosService([]domain.FundosDomain) *resterrors.RestErr
	CreateBalanceteService([]domain.BalanceteDomain) *resterrors.RestErr
	CreateExtratoService([]domain.ExtratoDomain) *resterrors.RestErr
	CreateInformacaoDiariaService([]domain.InformacaoDiariaDomain) *resterrors.RestErr
	CreatePerfilMensalService([]domain.PerfilMensalDomain) *resterrors.RestErr
	CreateCdaSelicService([]domain.CdaSelicDomain) *resterrors.RestErr
	CreateCdaBlc2Service([]domain.CdaBlc2Domain) *resterrors.RestErr
	CreateCdaBlc3Service([]domain.CdaBlc3Domain) *resterrors.RestErr
	CreateCdaBlc4Service([]domain.CdaBlc4Domain) *resterrors.RestErr
	CreateCdaBlc5Service([]domain.CdaBlc5Domain) *resterrors.RestErr
	CreateCdaBlc6Service([]domain.CdaBlc6Domain) *resterrors.RestErr
	CreateCdaBlc7Service([]domain.CdaBlc7Domain) *resterrors.RestErr
	CreateCdaBlc8Service([]domain.CdaBlc8Domain) *resterrors.RestErr
	CreateInformacaoComplementarFundoService([]domain.InformacoesFundoDomain) *resterrors.RestErr
	CreateInformacaoComplementarDivulgacaoService([]domain.InformacoesDivulgacaoDomain) *resterrors.RestErr
	CreateInformacaoComplementarCotistaService([]domain.InformacoesCotistaDomain) *resterrors.RestErr
	CreateInformacaoComplementarServicoPrestadoService([]domain.ServicoPrestadoDomain) *resterrors.RestErr
	CreateLaminaService([]domain.LaminaDomain) *resterrors.RestErr
	CreateLaminaCarteiraService([]domain.LaminaCarteiraDomain) *resterrors.RestErr
	CreateLaminaRentabilidadeAnoService([]domain.LaminaRentabilidadeAnoDomain) *resterrors.RestErr
	CreateLaminaRentabilidadeMesService([]domain.LaminaRentabilidadeMesDomain) *resterrors.RestErr
}
