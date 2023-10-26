package request

import "time"

type FundosCadastrosRequest struct {
	Id                    string    `json:"id"`
	AdministradorNome     string    `json:"administrador_nome" csv:"ADMIN,omitempty"`
	AuditorNome           string    `json:"auditor-nome" csv:"AUDITOR,omitempty"`
	CodigoCVM             int       `json:"codigo_cvm" csv:"CD_CVM,omitempty"`
	Classe                string    `json:"classe" csv:"CLASSE,omitempty"`
	ClasseAnbima          string    `json:"classe_anbima" csv:"CLASSE_ANBIMA,omitempty"`
	AdministradorCNPJ     string    `json:"administrador_cnpj" csv:"CNPJ_ADMIN,omitempty"`
	AuditorCNPJ           string    `json:"auditor_cnpj" csv:"CNPJ_AUDITOR,omitempty"`
	ControladorCNPJ       string    `json:"controlador_cnpj" csv:"CNPJ_CONTROLADOR,omitempty"`
	CustodianteCNPJ       string    `json:"custodiante_cnpj" csv:"CNPJ_CUSTODIANTE,omitempty"`
	FundoCNPJ             string    `json:"fundo_cnpj" csv:"CNPJ_FUNDO,omitempty"`
	Condominio            string    `json:"condominio" csv:"CONDOM,omitempty"`
	Controlador           string    `json:"controlador" csv:"CONTROLADOR,omitempty"`
	GestorCPFCNPJ         string    `json:"gestor_cpf_cnpj" csv:"CPF_CNPJ_GESTOR,omitempty"`
	Custodiante           string    `json:"custodiante" csv:"CUSTODIANTE,omitempty"`
	FundoNome             string    `json:"fundo_nome" csv:"DENOM_SOCIAL,omitempty"`
	Diretor               string    `json:"diretor" csv:"DIRETOR,omitempty"`
	DataCancelamento      time.Time `json:"data_cancelamento" csv:"DT_CANCEL,omitempty" time_format:"2006-01-02"`
	DataConstituicao      time.Time `json:"data_constituicao" csv:"DT_CONST,omitempty" time_format:"2006-01-02"`
	DataFim               time.Time `json:"data_fim" csv:"DT_FIM_EXERC,omitempty" time_format:"2006-01-02"`
	DataInicio            time.Time `json:"data_inicio" csv:"DT_INI_ATIV,omitempty" time_format:"2006-01-02"`
	DataInicioClasse      time.Time `json:"data_inicio_classe" csv:"DT_INI_CLASSE,omitempty" time_format:"2006-01-02"`
	DataInicioSocial      time.Time `json:"data_inicio_social" csv:"DT_INI_EXERC,omitempty" time_format:"2006-01-02"`
	DataInicioSituacao    time.Time `json:"data_inicio_situacao" csv:"DT_INI_SIT,omitempty" time_format:"2006-01-02"`
	DataPatrimonioLiquido time.Time `json:"data_patrimonio_liquido" csv:"DT_PATRIM_LIQ,omitempty" time_format:"2006-01-02"`
	DataRegistro          time.Time `json:"data_registro" csv:"DT_REG,omitempty" time_format:"2006-01-02"`
	Entidade              string    `json:"entidade" csv:"ENTID_INVEST,omitempty"`
	CotasPossui           string    `json:"cotas_possui" csv:"FUNDO_COTAS,omitempty"`
	Exclusivo             string    `json:"exclusivo" csv:"FUNDO_EXCLUSIVO,omitempty"`
	GestorNome            string    `json:"gestor_nome" csv:"GESTOR,omitempty"`
	InformacoesAdicionais string    `json:"informacoes_adicionais" csv:"INF_TAXA_ADM,omitempty"`
	InfoTaxaPerformance   string    `json:"info_taxa_performance" csv:"INF_TAXA_PERFM,omitempty"`
	Exterior100           string    `json:"exterior_100" csv:"INVEST_CEMPR_EXTER,omitempty"`
	PessoaFouPJ           string    `json:"pf_pj" csv:"PF_PJ_GESTOR,omitempty"`
	PublicoAlvo           string    `json:"publico_alvo" csv:"PUBLICO_ALVO,omitempty"`
	IndicadorDesempenho   string    `json:"indicador_desempenho" csv:"RENTAB_FUNDO,omitempty"`
	Situacao              string    `json:"situacao" csv:"SIT,omitempty"`
	TaxaAdm               float64   `json:"taxa_adm" csv:"TAXA_ADM,omitempty"`
	TaxaPerformance       float64   `json:"taxa_performance" csv:"TAXA_PERFM,omitempty"`
	TipoFundo             string    `json:"tipo_fundo" csv:"TP_FUNDO,omitempty"`
	TributacaoLongoPrazo  string    `json:"tributacao_longo_prazo" csv:"TRIB_LPRAZO,omitempty"`
	ValorPatriminioLiq    float64   `json:"valor_patriminio_liquido" csv:"VL_PATRIM_LIQ,omitempty"`
}
