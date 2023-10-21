package listener

import (
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
	"fmt"
)

func FundosSincroniszarListener(message []byte) *resterrors.RestErr {
	logger.Info("Init FundosSincroniszarListener", "sincronizar")
	var dat map[string]interface{}

	err := json.Unmarshal(message, &dat)
	if err != nil {
		logger.Error("json Unmarshal error", err, "listener")
	}
	logger.Info(
		fmt.Sprintf("json %v", dat),
		"sincronizar")

	logger.Info("Finish FundosSincroniszarListener", "sincronizar")

	return nil
}
