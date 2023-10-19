package controller

import (
	"api-fundos-investimentos/application/port/input"

	"github.com/gin-gonic/gin"
)

func NewFundosControllerInterface(
	serviceInterface input.FundosDomainService,
) FundosControllerInterface {
	return &fundosControllerInterface{
		service: serviceInterface,
	}
}

type FundosControllerInterface interface {
	SincronizarFundos(c *gin.Context)
}

type fundosControllerInterface struct {
	service input.FundosDomainService
}
