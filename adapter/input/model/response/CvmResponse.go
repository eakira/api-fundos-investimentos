package response

import "time"

type fundosExternoResponse struct {
	administradorNome     string    `form:"ADMIN" binding:"max=100"`
	auditorNome           string    `form:"AUDITOR" binding:"max=100"`
	codigoCVM             int16     `form:"CD_CVM" binding:"max=7"`
	classe                string    `form:"CLASSE" binding:"max=100"`
	classeAnbima          string    `form:"CLASSE_ANBIMA" binding:"max=100"`
	administradorCNPJ     string    `form:"CNPJ_ADMIN" binding:"max=20"`
	auditorCNPJ           string    `form:"CNPJ_AUDITOR" binding:"max=20"`
	controladorCNPJ       string    `form:"CNPJ_CONTROLADOR" binding:"max=20"`
	custodianteCNPJ       string    `form:"CNPJ_CUSTODIANTE" binding:"max=20"`
	fundoCNPJ             string    `form:"CNPJ_FUNDO" binding:"max=20"`
	condominio            string    `form:"CONDOM" binding:"max=100"`
	controlador           string    `form:"CONTROLADOR" binding:"max=100"`
	gestorCPFCNPJ         string    `form:"CPF_CNPJ_GESTOR" binding:"max=20"`
	custodiante           string    `form:"CUSTODIANTE" binding:"max=100"`
	fundoNome             string    `form:"DENOM_SOCIAL" binding:"max=100"`
	diretor               string    `form:"DIRETOR" binding:"max=100"`
	dataCancelamento      time.Time `form:"DT_CANCEL" time_format:"2006-01-02"`
	dataConstituicao      time.Time `form:"DT_CONST" time_format:"2006-01-02"`
	dataFim               time.Time `form:"DT_FIM_EXERC" time_format:"2006-01-02"`
	dataInicio            time.Time `form:"DT_INI_ATIV" time_format:"2006-01-02"`
	dataInicioClasse      time.Time `form:"DT_INI_CLASSE" time_format:"2006-01-02"`
	dataInicioSocial      time.Time `form:"DT_INI_EXERC" time_format:"2006-01-02"`
	dataInicioSituacao    time.Time `form:"DT_INI_SIT" time_format:"2006-01-02"`
	dataPatrimonioLiquido time.Time `form:"DT_PATRIM_LIQ" time_format:"2006-01-02"`
	dataRegistro          time.Time `form:"DT_REG" time_format:"2006-01-02"`
	entidade              string    `form:"ENTID_INVEST" binding:"max=1"`
	cotasPossui           string    `form:"FUNDO_COTAS" binding:"max=1"`
	exclusivo             string    `form:"FUNDO_EXCLUSIVO" binding:"max=1"`
	gestorNome            string    `form:"GESTOR" binding:"max=100"`
	informacoesAdicionais string    `form:"INF_TAXA_ADM" binding:"max=400"`
	InfoTaxaPerformance   string    `form:"INF_TAXA_PERFM" binding:"max=400"`
	exterior100           string    `form:"INVEST_CEMPR_EXTER" binding:"max=1"`
	PessoaFOuPJ           string    `form:"PF_PJ_GESTOR" binding:"max=2"`
	publicoAlvo           string    `form:"PUBLICO_ALVO" binding:"max=15"`
	indicadorDesempenho   string    `form:"RENTAB_FUNDO" binding:"max=100"`
	situacao              string    `form:"SIT" binding:"max=100"`
	taxaAdm               float64   `form:"TAXA_ADM"`
	taxaPerformance       float64   `form:"TAXA_PERFM"`
	tipoFundo             string    `form:"TP_FUNDO" binding:"max=20"`
	tributacaoLongoPrazo  string    `form:"TRIB_LPRAZO" binding:"max=3"`
	valorPatriminioLiq    float64   `form:"VL_PATRIM_LIQ"`
}
