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
		dados := []request.CdaBlc1Request{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.CdaBlc1Domain{}
		copier.Copy(domain, dados)
		fs.CreateCdaBlc1Service(*domain)

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
		fs.CreateInformaComplementarFundoService(*domain)

	case "informacao-complementar-divulgacao":
		dados := []request.InformacoesDivulgacaoRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacoesDivulgacaoDomain{}
		copier.Copy(domain, dados)
		fs.CreateInformaComplementarDivulgacaoService(*domain)

	case "informacao-complementar-cotista":
		dados := []request.InformacoesCotistaResquest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.InformacoesCotistaDomain{}
		copier.Copy(domain, dados)
		fs.CreateInformaComplementarCotistaService(*domain)

	case "informacao-complementar-servico-prestado":
		dados := []request.ServicoPrestadoResquest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.ServicoPrestadoDomain{}
		copier.Copy(domain, dados)
		fs.CreateInformaComplementarServicoPrestadoService(*domain)

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

	case "lamina":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "perfil-mensal":
		//		files = getFilesName(env.GetConfigCvmArquivosPerfilMensal())

	}
}
