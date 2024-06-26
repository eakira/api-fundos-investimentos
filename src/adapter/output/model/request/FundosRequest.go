package request

type FundosRequest struct {
	AdministradorNome     string  `json:"ADMIN,omitempty"`
	AuditorNome           string  `json:"AUDITOR,omitempty"`
	CodigoCVM             Integer `json:"CD_CVM,omitempty"`
	Classe                string  `json:"CLASSE,omitempty"`
	ClasseAnbima          string  `json:"CLASSE_ANBIMA,omitempty"`
	AdministradorCNPJ     Cnpj    `json:"CNPJ_ADMIN,omitempty"`
	AuditorCNPJ           Cnpj    `json:"CNPJ_AUDITOR,omitempty"`
	ControladorCNPJ       Cnpj    `json:"CNPJ_CONTROLADOR,omitempty"`
	CustodianteCNPJ       Cnpj    `json:"CNPJ_CUSTODIANTE,omitempty"`
	FundoCNPJ             Cnpj    `json:"CNPJ_FUNDO,omitempty"`
	Condominio            string  `json:"CONDOM,omitempty"`
	Controlador           string  `json:"CONTROLADOR,omitempty"`
	GestorCPFCNPJ         string  `json:"CPF_CNPJ_GESTOR,omitempty"`
	Custodiante           string  `json:"CUSTODIANTE,omitempty"`
	FundoNome             string  `json:"DENOM_SOCIAL,omitempty"`
	Diretor               string  `json:"DIRETOR,omitempty"`
	DataCancelamento      Date    `json:"DT_CANCEL,omitempty"`
	DataConstituicao      Date    `json:"DT_CONST,omitempty"`
	DataFim               Date    `json:"DT_FIM_EXERC,omitempty"`
	DataInicio            Date    `json:"DT_INI_ATIV,omitempty"`
	DataInicioClasse      Date    `json:"DT_INI_CLASSE,omitempty"`
	DataInicioSocial      Date    `json:"DT_INI_EXERC,omitempty"`
	DataInicioSituacao    Date    `json:"DT_INI_SIT,omitempty"`
	DataPatrimonioLiquido Date    `json:"DT_PATRIM_LIQ,omitempty"`
	DataRegistro          Date    `json:"DT_REG,omitempty"`
	Entidade              string  `json:"ENTID_INVEST,omitempty"`
	CotasPossui           string  `json:"FUNDO_COTAS,omitempty"`
	Exclusivo             string  `json:"FUNDO_EXCLUSIVO,omitempty"`
	GestorNome            string  `json:"GESTOR,omitempty"`
	InformacoesAdicionais string  `json:"INF_TAXA_ADM,omitempty"`
	InfoTaxaPerformance   string  `json:"INF_TAXA_PERFM,omitempty"`
	Exterior100           string  `json:"INVEST_CEMPR_EXTER,omitempty"`
	PessoaFouPJ           string  `json:"PF_PJ_GESTOR,omitempty"`
	PublicoAlvo           string  `json:"PUBLICO_ALVO,omitempty"`
	IndicadorDesempenho   string  `json:"RENTAB_FUNDO,omitempty"`
	Situacao              string  `json:"SIT,omitempty"`
	TaxaAdm               string  `json:"TAXA_ADM,omitempty"`
	TaxaPerformance       string  `json:"TAXA_PERFM,omitempty"`
	TipoFundo             string  `json:"TP_FUNDO,omitempty"`
	TributacaoLongoPrazo  string  `json:"TRIB_LPRAZO,omitempty"`
	ValorPatriminioLiq    Decimal `json:"VL_PATRIM_LIQ,omitempty"`
}
