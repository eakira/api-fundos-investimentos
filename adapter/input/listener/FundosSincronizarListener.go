package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
	"fmt"
)

func FundosSincroniszarListener(
	message []byte,
	controller controller.FundosControllerInterface) *resterrors.RestErr {
	logger.Info("Init FundosSincroniszarListener", "sincronizar")
	dados := &request.FundosQueueSincronizarRequest{}

	err := json.Unmarshal(message, &dados)
	if err != nil {
		logger.Error("json Unmarshal error", err, "listener")
	}
	logger.Info(
		fmt.Sprintf("json %v", dados),
		"sincronizar")

	controller.DownloadArquivosCVMController(dados.Data[0])

	logger.Info("Finish FundosSincroniszarListener", "sincronizar")

	return nil
}
