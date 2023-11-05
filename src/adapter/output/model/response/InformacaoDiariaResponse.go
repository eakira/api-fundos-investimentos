package response

type InformacaoDiariaResponse struct {
	FundoCnpj              string `json:"CNPJ_FUNDO,omitempty"`
	CaptacaoDia            string `json:"CAPTC_DIA,omitempty"`
	DataCompetencia        string `json:"DT_COMPTC,omitempty"`
	ResgateDia             string `json:"RESG_DIA,omitempty"`
	TipoFundo              string `json:"TP_FUNDO,omitempty"`
	ValorPatrimonioLiquido string `json:"VL_PATRIM_LIQ,omitempty"`
	CotaValor              string `json:"VL_QUOTA,omitempty"`
	QtdCotista             string `json:"NR_COTST,omitempty"`
	ValorTotal             string `json:"VL_TOTAL,omitempty"`
}
