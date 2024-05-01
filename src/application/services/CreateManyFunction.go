package services

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"

	"github.com/jinzhu/copier"
)

// TODO: REFATORAR
func (fs *fundosDomainService) CreateMany(
	data []byte,
) *resterrors.RestErr {
	logger.Info("Init FundosPersistenciaDadosListener", "sincronizarFundos")
	mapa := []map[string]string{}

	err := json.Unmarshal(data, &mapa)
	if err != nil {
		logger.Error("json Unmarshal error", err, "listener")
	}

	switch mapa[0]["collection"] {
	case "cadastros":
		dados := []request.FundosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.FundosDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateFundosService(*domain)
		if erro != nil {
			return erro
		}
	case "balancete":
		dados := []request.BalanceteRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.BalanceteDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateBalanceteService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-blc-1":
		dados := []request.CdaSelicRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaSelicDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaSelicService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-blc-2":
		dados := []request.CdaFundosInvestimentosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaFundosInvestimentosDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaFundosInvestimentosService(*domain)

	case "cda-blc-3":
		dados := []request.CdaSwapRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaSwapDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaSwapService(*domain)

	case "cda-blc-4":
		dados := []request.CdaDemaisAtivosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaDemaisAtivosDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaDemaisAtivosService(*domain)

	case "cda-blc-5":
		dados := []request.CdaDepositoAPrazoOutrosAtivosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaDepositoAPrazoOutrosAtivosDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaDepositoAPrazoOutrosAtivosService(*domain)

	case "cda-blc-6":
		dados := []request.CdaAgroCreditoPrivadoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaAgroCreditoPrivadoDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaAgroCreditoPrivadoService(*domain)

	case "cda-blc-7":
		dados := []request.CdaInvestimentosExteriorRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaInvestimentosExteriorDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaInvestimentosExteriorService(*domain)

	case "cda-blc-8":
		dados := []request.CdaDemaisAtivosNaoCodificadosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaDemaisAtivosNaoCodificadosDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaDemaisAtivosNaoCodificadosService(*domain)

	case "cda-confidencial":
		dados := []request.CdaConfidencialRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaConfidencialDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaConfidencialService(*domain)

	case "cda-fiim":
		dados := []request.CdaFiimRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaFiimDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaFiimService(*domain)

	case "cda-fiim-confidencial":
		dados := []request.CdaFiimConfidencialRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaFiimConfidencialDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaFiimConfidencialidade(*domain)

	case "cda-patrimonio-liquido":
		dados := []request.CdaPatrimonioLiquidoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaPatrimonioLiquidoDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaPatrominioLiquido(*domain)

	case "informacao-complementar":
		dados := []request.InformacoesFundoResquest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacoesFundoDomain{}
		copier.Copy(domain, dados)
		fs.CreateInformacaoComplementarFundoService(*domain)

	case "informacao-complementar-divulgacao":
		dados := []request.InformacoesDivulgacaoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacoesDivulgacaoDomain{}
		copier.Copy(domain, dados)
		fs.CreateInformacaoComplementarDivulgacaoService(*domain)

	case "informacao-complementar-cotista":
		dados := []request.InformacoesCotistaResquest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacoesCotistaDomain{}
		copier.Copy(domain, dados)
		fs.CreateInformacaoComplementarCotistaService(*domain)

	case "informacao-complementar-servico-prestado":
		dados := []request.ServicoPrestadoResquest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.ServicoPrestadoDomain{}
		copier.Copy(domain, dados)
		fs.CreateInformacaoComplementarServicoPrestadoService(*domain)

	case "lamina":
		dados := []request.LaminaRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.LaminaDomain{}
		copier.Copy(domain, dados)
		fs.CreateLaminaService(*domain)

	case "lamina-carteira":
		dados := []request.LaminaCarteiraRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.LaminaCarteiraDomain{}
		copier.Copy(domain, dados)
		fs.CreateLaminaCarteiraService(*domain)

	case "lamina-rentabilidade-ano":
		dados := []request.LaminaRentabilidadeAnoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.LaminaRentabilidadeAnoDomain{}
		copier.Copy(domain, dados)
		fs.CreateLaminaRentabilidadeAnoService(*domain)

	case "lamina-rentabilidade-mes":
		dados := []request.LaminaRentabilidadeMesRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.LaminaRentabilidadeMesDomain{}
		copier.Copy(domain, dados)
		fs.CreateLaminaRentabilidadeMesService(*domain)

	case "extrato":
		dados := []request.ExtratoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.ExtratoDomain{}
		copier.Copy(domain, dados)
		fs.CreateExtratoService(*domain)

	case "informacao-diaria":
		dados := []request.InformacaoDiariaRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacaoDiariaDomain{}
		copier.Copy(domain, dados)
		fs.CreateInformacaoDiariaService(*domain)

	case "perfil-mensal":
		dados := []request.PerfilMensalRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.PerfilMensalDomain{}
		copier.Copy(domain, dados)
		fs.CreatePerfilMensalService(*domain)

	}

	return nil
}
