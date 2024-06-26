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
		erro := fs.CreateCdaFundosInvestimentosService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-blc-3":
		dados := []request.CdaSwapRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaSwapDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaSwapService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-blc-4":
		dados := []request.CdaDemaisAtivosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaDemaisAtivosDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaDemaisAtivosService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-blc-5":
		dados := []request.CdaDepositoAPrazoOutrosAtivosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaDepositoAPrazoOutrosAtivosDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaDepositoAPrazoOutrosAtivosService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-blc-6":
		dados := []request.CdaAgroCreditoPrivadoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaAgroCreditoPrivadoDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaAgroCreditoPrivadoService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-blc-7":
		dados := []request.CdaInvestimentosExteriorRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaInvestimentosExteriorDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaInvestimentosExteriorService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-blc-8":
		dados := []request.CdaDemaisAtivosNaoCodificadosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaDemaisAtivosNaoCodificadosDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaDemaisAtivosNaoCodificadosService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-confidencial":
		dados := []request.CdaConfidencialRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaConfidencialDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaConfidencialService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-fiim":
		dados := []request.CdaFiimRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaFiimDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaFiimService(*domain)
		if erro != nil {
			return erro
		}

	case "cda-fiim-confidencial":
		dados := []request.CdaFiimConfidencialRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaFiimConfidencialDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaFiimConfidencialidade(*domain)
		if erro != nil {
			return erro
		}

	case "cda-patrimonio-liquido":
		dados := []request.CdaPatrimonioLiquidoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaPatrimonioLiquidoDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateCdaPatrominioLiquido(*domain)
		if erro != nil {
			return erro
		}

	case "informacao-complementar":
		dados := []request.InformacoesFundoResquest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacoesFundoDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateInformacaoComplementarFundoService(*domain)
		if erro != nil {
			return erro
		}

	case "informacao-complementar-divulgacao":
		dados := []request.InformacoesDivulgacaoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacoesDivulgacaoDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateInformacaoComplementarDivulgacaoService(*domain)
		if erro != nil {
			return erro
		}

	case "informacao-complementar-cotista":
		dados := []request.InformacoesCotistaResquest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacoesCotistaDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateInformacaoComplementarCotistaService(*domain)
		if erro != nil {
			return erro
		}

	case "informacao-complementar-servico-prestado":
		dados := []request.ServicoPrestadoResquest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.ServicoPrestadoDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateInformacaoComplementarServicoPrestadoService(*domain)
		if erro != nil {
			return erro
		}

	case "lamina":
		dados := []request.LaminaRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.LaminaDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateLaminaService(*domain)
		if erro != nil {
			return erro
		}

	case "lamina-carteira":
		dados := []request.LaminaCarteiraRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.LaminaCarteiraDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateLaminaCarteiraService(*domain)
		if erro != nil {
			return erro
		}

	case "lamina-rentabilidade-ano":
		dados := []request.LaminaRentabilidadeAnoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.LaminaRentabilidadeAnoDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateLaminaRentabilidadeAnoService(*domain)
		if erro != nil {
			return erro
		}

	case "lamina-rentabilidade-mes":
		dados := []request.LaminaRentabilidadeMesRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.LaminaRentabilidadeMesDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateLaminaRentabilidadeMesService(*domain)
		if erro != nil {
			return erro
		}

	case "extrato":
		dados := []request.ExtratoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.ExtratoDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateExtratoService(*domain)
		if erro != nil {
			return erro
		}

	case "informacao-diaria":
		dados := []request.InformacaoDiariaRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacaoDiariaDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreateInformacaoDiariaService(*domain)
		if erro != nil {
			return erro
		}

	case "perfil-mensal":
		dados := []request.PerfilMensalRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.PerfilMensalDomain{}
		copier.Copy(domain, dados)
		erro := fs.CreatePerfilMensalService(*domain)
		if erro != nil {
			return erro
		}

	}

	return nil
}
