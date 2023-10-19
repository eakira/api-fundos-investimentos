package services

import (
	"api-fundos-investimentos/application/port/input"
	"api-fundos-investimentos/application/port/output"
)

func NewFundosDomainService(
	fundosRepository output.FundosPort, fundosHttp output.FundosExternoPort) input.FundosDomainService {
	return &fundosDomainService{
		fundosRepository,
		fundosHttp,
	}
}

type fundosDomainService struct {
	repository output.FundosPort
	http       output.FundosExternoPort
}
