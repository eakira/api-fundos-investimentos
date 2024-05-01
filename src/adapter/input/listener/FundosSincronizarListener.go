package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
)

func FundosSincronizarListener(
	message []byte,
	controller controller.FundosControllerInterface,
) *resterrors.RestErr {
	logger.Info("Init FundosSincronizarListener", "sincronizarFundos")
	dados := request.FundosCvmArquivosQueueRequest{}

	err := json.Unmarshal(message, &dados)
	if err != nil {
		logger.Error("json Unmarshal error", err, "sincronizarFundos")
	}

	controller.DownloadArquivosCVMController(dados)

	logger.Info("Finish FundosSincronizarListener", "sincronizarFundos")

	return nil
}
