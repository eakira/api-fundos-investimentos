package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
	"fmt"
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
	logger.Info(
		fmt.Sprintf("json %s", dados),
		"sincronizarFundos")

	controller.DownloadArquivosCVMController(dados)

	logger.Info("Finish FundosSincronizarListener", "sincronizarFundos")

	return nil
}
