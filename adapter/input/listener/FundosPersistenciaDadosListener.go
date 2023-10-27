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
	mapa := map[string]string{}

	err := json.Unmarshal(message, &mapa)
	if err != nil {
		logger.Error("json Unmarshal error", err, "listener")
	}

	switch mapa["collection"] {
	case "cadastros":
		dados := request.FundosCadastrosRequest{}
		json.Unmarshal(message, &dados)
		controller.CreateFundosController(dados)
		fmt.Println(dados)
		fmt.Println(mapa)
		panic("teste")

	case "balancete":

	case "cda":
		//		files = getFilesName(env.GetConfigCvmCda())
		//		histfiles := getFilesName(env.GetConfigCvmCdaHist())
		//		files = append(files, histfiles...)

	case "informacoes-complementares":
		//		files = getFilesName(env.GetConfigCvmInformacoesComplementares())

	case "extrato":

	case "informacao-diaria":

	case "lamina":
		//		files = getFilesName(env.GetConfigCvmArquivosLaminas())
		//		histfiles := getFilesName(env.GetConfigCvmArquivosLaminasHist())
		//		files = append(files, histfiles...)

	case "perfil-mensal":
		//		files = getFilesName(env.GetConfigCvmArquivosPerfilMensal())

	}

	logger.Info("Finish FundosProcessarArquivosListener", "sincronizar")

	return nil
}
