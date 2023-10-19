package controller

import "api-fundos-investimentos/application/port/input"

func NewFundosControllerInterface(
	serviceInterface input.FundosDomainService,
) FundosControllerInterface {
	return &fundosControllerInterface{
		service: serviceInterface,
	}
}

type FundosControllerInterface interface {
}

type fundosControllerInterface struct {
	service input.FundosDomainService
}
