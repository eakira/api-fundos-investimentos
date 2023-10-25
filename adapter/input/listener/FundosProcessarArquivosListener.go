package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
	"fmt"
)

func FundosProcessarArquivosListener(
	message []byte,
	controller controller.FundosControllerInterface,
) *resterrors.RestErr {
	logger.Info("Init FundosProcessarArquivosListener", "sincronizar")
	dados := request.FundosCvmArquivosQueueRequest{}

	err := json.Unmarshal(message, &dados)
	if err != nil {
		logger.Error("json Unmarshal error", err, "listener")
	}
	logger.Info(
		fmt.Sprintf("json %s", dados),
		"sincronizar")

	controller.DownloadArquivosCVMController(dados)

	logger.Info("Finish FundosProcessarArquivosListener", "sincronizar")

	return nil
}
