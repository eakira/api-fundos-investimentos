package services

import (
	"api-fundos-investimentos/application/port/input"
	"api-fundos-investimentos/application/port/output"
)

func NewFundosDomainService(
	fundosRepository output.FundosPort) input.FundosDomainService {
	return &fundosDomainService{
		fundosRepository,
	}
}

type fundosDomainService struct {
	repository output.FundosPort
}
