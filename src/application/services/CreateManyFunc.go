package services

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"encoding/json"

	"github.com/jinzhu/copier"
)

func CreateMany(
	fs *fundosDomainService,
	data []byte,
) {
	logger.Info("Init FundosPersistenciaDadosListener", "sincronizar")
	mapa := []map[string]string{}

	err := json.Unmarshal(data, &mapa)
	if err != nil {
		logger.Error("json Unmarshal error", err, "listener")
	}

	switch mapa[0]["collection"] {
	case "cadastros":
		dados := []request.FundosCadastrosRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.FundosDomain{}
		copier.Copy(domain, dados)
		fs.CreateFundosService(*domain)

	case "balancete":
		dados := []request.BalanceteRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.BalanceteDomain{}
		copier.Copy(domain, dados)
		fs.CreateBalanceteService(*domain)

	case "cda-blc-1":
		dados := []request.CdaSelicRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaSelicDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaSelicService(*domain)

	case "cda-blc-2":
		dados := []request.CdaBlc2Request{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaBlc2Domain{}
		copier.Copy(domain, dados)
		fs.CreateCdaBlc2Service(*domain)

	case "cda-blc-3":
		dados := []request.CdaBlc3Request{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaBlc3Domain{}
		copier.Copy(domain, dados)
		fs.CreateCdaBlc3Service(*domain)

	case "cda-blc-4":
		dados := []request.CdaBlc4Request{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaBlc4Domain{}
		copier.Copy(domain, dados)
		fs.CreateCdaBlc4Service(*domain)

	case "cda-blc-5":
		dados := []request.CdaBlc5Request{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaBlc5Domain{}
		copier.Copy(domain, dados)
		fs.CreateCdaBlc5Service(*domain)

	case "cda-blc-6":
		dados := []request.CdaBlc6Request{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaBlc6Domain{}
		copier.Copy(domain, dados)
		fs.CreateCdaBlc6Service(*domain)

	case "cda-blc-7":
		dados := []request.CdaBlc7Request{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaBlc7Domain{}
		copier.Copy(domain, dados)
		fs.CreateCdaBlc7Service(*domain)

	case "cda-blc-8":
		dados := []request.CdaBlc8Request{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaBlc8Domain{}
		copier.Copy(domain, dados)
		fs.CreateCdaBlc8Service(*domain)

	case "cda-confidencial":
		dados := []request.CdaConfidencial{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaConfidencialDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaConfidencialService(*domain)

	case "cda-fiim":
		dados := []request.CdaFiim{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaFiimDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaFiimService(*domain)

	case "cda-fiim-confidencial":
		dados := []request.CdaFiimConfidencial{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaFiimConfidencialDomain{}
		copier.Copy(domain, dados)
		fs.CreateCdaFiimConfidencialidade(*domain)

	case "cda-patrimonio-liquido":
		dados := []request.CdaPatrimonioLiquido{}
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
}
