package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
)

func FundosPersistenciaDadosListener(
	message []byte,
	controller controller.FundosControllerInterface,
) *resterrors.RestErr {
	logger.Info("Init FundosPersistenciaDadosListener", "sincronizarFundos")
	mapa := []map[string]string{}

	err := json.Unmarshal(message, &mapa)
	if err != nil {
		logger.Error("json Unmarshal error", err, "sincronizarFundos")
	}

	switch mapa[0]["collection"] {
	case "cadastros":
		dados := []request.FundosCadastrosRequest{}
		json.Unmarshal(message, &dados)
		controller.CreateFundosController(dados)

	case "balancete":
		dados := []request.BalanceteRequest{}
		json.Unmarshal(message, &dados)
		controller.CreateBalanceteController(dados)

	case "cda":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "informacoes-complementares":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "extrato":
		dados := []request.ExtratoRequest{}
		json.Unmarshal(message, &dados)
		controller.CreateExtratoController(dados)

	case "informacao-diaria":
		dados := []request.InformacaoDiariaRequest{}
		json.Unmarshal(message, &dados)
		controller.CreateInformacaoDiariaController(dados)

	case "lamina":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "perfil-mensal":
		//		files = getFilesName(env.GetConfigCvmArquivosPerfilMensal())

	}

	logger.Info("Finish FundosProcessarArquivosListener", "sincronizarFundos")

	return nil
}
