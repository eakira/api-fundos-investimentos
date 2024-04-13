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
	logger.Info("Init FundosProcessarArquivosListener", "sincronizarFundos")
	dados := request.FundosCvmArquivosQueueRequest{}

	err := json.Unmarshal(message, &dados)
	if err != nil {
		logger.Error("json Unmarshal error", err, "sincronizarFundos")
	}
	logger.Info(
		fmt.Sprintf("json %s", dados),
		"sincronizarFundos")

	controller.ProcessarArquivosCVMController(dados)

	logger.Info("Finish FundosProcessarArquivosListener", "sincronizarFundos")

	return nil
}
