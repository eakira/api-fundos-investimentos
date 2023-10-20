package output

import "api-fundos-investimentos/configuration/resterrors"

type FundosPort interface {
}

type FundosExternoPort interface {
}

type FundosQueuePort interface {
	Produce(string, string) *resterrors.RestErr
}
