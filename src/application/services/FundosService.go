package services

import (
	"api-fundos-investimentos/application/port/input"
	"api-fundos-investimentos/application/port/output"
)

func NewFundosDomainService(
	fundosRepository output.FundosPort,
	queue output.FundosQueuePort,
	externo output.FundosExternoPort,
) input.FundosDomainService {
	return &fundosDomainService{
		fundosRepository,
		queue,
		externo,
	}
}

type fundosDomainService struct {
	repository output.FundosPort
	queue      output.FundosQueuePort
	externo    output.FundosExternoPort
}
