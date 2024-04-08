package request

import "time"

type CdaBlc1Request struct {
	CodigoConta              string    `json:"CD_CONTA_BALCTE,omitempty"`    // Código da conta balancete
	FundoCnpj                string    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo
	DataCompetencia          string    `json:"DT_COMPTC,omitempty"`          // Data de competência
	PlanoContabil            string    `json:"PLANO_CONTA_BALCTE,omitempty"` // Plano contábil
	SaldoConta               string    `json:"VL_SALDO_BALCTE,omitempty"`    // Saldo da conta balancete
	CdIsin                   string    `json:"CD_ISIN,omitempty"`            // Código ISIN (International Securities Identification Number)
	CdSelic                  string    `json:"CD_SELIC,omitempty"`           // Código SELIC
	DenominacaoSocial        string    `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social
	DataEmissao              time.Time `json:"DT_EMISSAO,omitempty"`         // Data de emissão (AAAA-MM-DD)
	DataVencimento           time.Time `json:"DT_VENC,omitempty"`            // Data de vencimento (AAAA-MM-DD)
	DataCompetenciaDoc       string    `json:"DT_COMPTC_DOC,omitempty"`      // Data de competência do documento
	PrazoConfidencialidade   time.Time `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	EmissorLigado            string    `json:"EMISSOR_LIGADO,omitempty"`     // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	QuantidadeAquisicoes     float64   `json:"QT_AQUIS_NEGOC,omitempty"`     // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal   float64   `json:"QT_POS_FINAL,omitempty"`       // Quantidade da posição final
	QuantidadeVendas         float64   `json:"QT_VENDA_NEGOC,omitempty"`     // Quantidade de vendas dos negócios realizados no mês
	TipoAplicacao            string    `json:"TP_APLIC,omitempty"`           // Tipo de aplicação
	TipoAtivo                string    `json:"TP_ATIVO,omitempty"`           // Tipo de ativo
	TipoFundo                string    `json:"TP_FUNDO,omitempty"`           // Tipo de fundo
	TipoNegociacao           string    `json:"TP_NEGOC,omitempty"`           // Tipo de negociação (Para negociação/Mantido até o vencimento)
	TipoTituloPublico        string    `json:"TP_TITPUB,omitempty"`          // Tipo de título público
	ValorAquisicoes          float64   `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal   float64   `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal float64   `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas              float64   `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}
