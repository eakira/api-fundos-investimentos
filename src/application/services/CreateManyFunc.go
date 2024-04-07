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

	case "cda":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "informacoes-complementares":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "extrato":
		dados := []request.ExtratoRequest{}
		json.Unmarshal(data, &dados)

	case "informacao-diaria":
		dados := []request.InformacaoDiariaRequest{}
		json.Unmarshal(data, &dados)
		domain := &[]domain.ExtratoDomain{}
		copier.Copy(domain, dados)
		fs.CreateExtratoService(*domain)

	case "lamina":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "perfil-mensal":
		//		files = getFilesName(env.GetConfigCvmArquivosPerfilMensal())

	}
}
