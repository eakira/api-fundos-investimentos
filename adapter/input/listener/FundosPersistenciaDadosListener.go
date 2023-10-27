package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
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
	dados := map[string]string{}

	err := json.Unmarshal(message, &dados)
	if err != nil {
		logger.Error("json Unmarshal error", err, "listener")
	}

	fmt.Println(dados)

	logger.Info("Finish FundosProcessarArquivosListener", "sincronizar")

	return nil
}
