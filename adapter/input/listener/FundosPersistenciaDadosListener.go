package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
	"fmt"
)

func FundosPersistenciaDadosListener(
	message []byte,
	controller controller.FundosControllerInterface,
) *resterrors.RestErr {
	logger.Info("Init FundosPersistenciaDadosListener", "sincronizar")
	mapa := []map[string]string{}

	err := json.Unmarshal(message, &mapa)
	if err != nil {
		logger.Error("json Unmarshal error", err, "listener")
	}

	switch mapa[0]["collection"] {
	case "cadastros":
		dados := []request.FundosCadastrosRequest{}
		json.Unmarshal(message, &dados)
		fmt.Println(dados)
		controller.CreateFundosController(dados)

	case "balancete":
		//		dados := []request.BalanceteRequest{}
		var data []interface{}
		json.Unmarshal(message, data)
		fmt.Println(data)
		panic("Aqui")
		//		controller.CreateBalanceteController(dados)

	case "cda":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "informacoes-complementares":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "extrato":
		//		dados := []request.ExtratoRequest{}
		var data []interface{}
		json.Unmarshal(message, data)
		panic("Aqui")
		//		controller.CreateExtratoController(dados)

	case "informacao-diaria":
		// dados := []request.InformacaoDiariaRequest{}
		var data []interface{}
		json.Unmarshal(message, data)
		panic("Aqui")
		//		controller.CreateInformacaoDiariaController(dados)

	case "lamina":
		//		São vários arquivos precisa verificar quais arquivos vou usar

	case "perfil-mensal":
		//		files = getFilesName(env.GetConfigCvmArquivosPerfilMensal())

	}

	logger.Info("Finish FundosProcessarArquivosListener", "sincronizar")

	return nil
}
