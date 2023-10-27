package request

type BalanceteRequest struct {
	CodigoConta     string `json:"CD_CONTA_BALCTE,omitempty"`
	FundoCnpj       string `json:"CNPJ_FUNDO,omitempty"`
	DataCompetencia string `json:"DT_COMPTC,omitempty"`
	PlanoContabil   string `json:"PLANO_CONTA_BALCTE,omitempty"`
	SaldoConta      string `json:"VL_SALDO_BALCTE,omitempty"`
}
