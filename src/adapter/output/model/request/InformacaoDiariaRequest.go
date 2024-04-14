package request

type InformacaoDiariaRequest struct {
	FundoCnpj              Cnpj    `json:"CNPJ_FUNDO,omitempty"`
	CaptacaoDia            Decimal `json:"CAPTC_DIA,omitempty"`
	DataCompetencia        Date    `json:"DT_COMPTC,omitempty"`
	ResgateDia             Decimal `json:"RESG_DIA,omitempty"`
	TipoFundo              string  `json:"TP_FUNDO,omitempty"`
	ValorPatrimonioLiquido Decimal `json:"VL_PATRIM_LIQ,omitempty"`
	CotaValor              Decimal `json:"VL_QUOTA,omitempty"`
	QtdCotista             Integer `json:"NR_COTST,omitempty"`
	ValorTotal             Decimal `json:"VL_TOTAL,omitempty"`
}
