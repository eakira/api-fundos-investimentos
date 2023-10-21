package listener

import (
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func FundosSincroniszarListener(message string) *resterrors.RestErr {
	logger.Info("Init FundosSincroniszarListener", "sincronizar")

	logger.Info("Finish FundosSincroniszarListener", "sincronizar")

	return nil
}
