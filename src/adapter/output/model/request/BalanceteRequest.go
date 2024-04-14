package request

type BalanceteRequest struct {
	CodigoConta     Integer `json:"CD_CONTA_BALCTE,omitempty"`
	FundoCnpj       Cnpj    `json:"CNPJ_FUNDO,omitempty"`
	DataCompetencia Date    `json:"DT_COMPTC,omitempty"`
	PlanoContabil   string  `json:"PLANO_CONTA_BALCTE,omitempty"`
	SaldoConta      Decimal `json:"VL_SALDO_BALCTE,omitempty"`
}
