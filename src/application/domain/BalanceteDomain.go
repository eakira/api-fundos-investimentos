package domain

import "time"

type BalanceteDomain struct {
	Id              string
	CodigoConta     int
	FundoCnpj       string
	DataCompetencia time.Time
	PlanoContabil   string
	SaldoConta      float64
}
