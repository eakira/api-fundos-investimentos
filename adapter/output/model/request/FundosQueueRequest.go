package request

import "time"

type FundosQueueSincronizarRequest struct {
	Queue string   `json:"queue"`
	Data  []string `json:"data"`
}

type FundosCvmArquivosQueueRequest struct {
	Id          string    `json:"id"`
	Endereco    string    `json:"endereco"`
	TipoArquivo string    `json:"tipo-arquivo"`
	Referencia  string    `json:"referencia"`
	Status      string    `json:"status"`
	Download    string    `json:"download"`
	Processado  string    `json:"processado"`
	CreatedAt   time.Time `json:"created-at"`
	UpdateAt    time.Time `json:"update-at"`
}

type FundosCadastrosRequest struct {
	Id                    string    `json:"id"`
	AdministradorNome     string    `json:"administrador_nome"`
	AuditorNome           string    `json:"auditor-nome"`
	CodigoCVM             int16     `json:"codigo_cvm"`
	Classe                string    `json:"classe"`
	ClasseAnbima          string    `json:"classe_anbima"`
	AdministradorCNPJ     string    `json:"administrador_cnpj"`
	AuditorCNPJ           string    `json:"auditor_cnpj"`
	ControladorCNPJ       string    `json:"controlador_cnpj"`
	CustodianteCNPJ       string    `json:"custodiante_cnpj"`
	FundoCNPJ             string    `json:"fundo_cnpj"`
	Condominio            string    `json:"condominio"`
	Controlador           string    `json:"controlador"`
	GestorCPFCNPJ         string    `json:"gestor_cpf_cnpj"`
	Custodiante           string    `json:"custodiante"`
	FundoNome             string    `json:"fundo_nome"`
	Diretor               string    `json:"diretor"`
	DataCancelamento      time.Time `json:"data_cancelamento"`
	DataConstituicao      time.Time `json:"data_constituicao"`
	DataFim               time.Time `json:"data_fim"`
	DataInicio            time.Time `json:"data_inicio"`
	DataInicioClasse      time.Time `json:"data_inicio_classe"`
	DataInicioSocial      time.Time `json:"data_inicio_social"`
	DataInicioSituacao    time.Time `json:"data_inicio_situacao"`
	DataPatrimonioLiquido time.Time `json:"data_patrimonio_liquido"`
	DataRegistro          time.Time `json:"data_registro"`
	Entidade              string    `json:"entidade"`
	CotasPossui           string    `json:"cotas_possui"`
	Exclusivo             string    `json:"exclusivo"`
	GestorNome            string    `json:"gestor_nome"`
	InformacoesAdicionais string    `json:"informacoes_adicionais"`
	InfoTaxaPerformance   string    `json:"info_taxa_performance"`
	Exterior100           string    `json:"exterior_100"`
	PessoaFouPJ           string    `json:"pf_pj"`
	PublicoAlvo           string    `json:"publico_alvo"`
	IndicadorDesempenho   string    `json:"indicador_desempenho"`
	Situacao              string    `json:"situacao"`
	TaxaAdm               float64   `json:"taxa_adm"`
	TaxaPerformance       float64   `json:"taxa_performance"`
	TipoFundo             string    `json:"tipo_fundo"`
	TributacaoLongoPrazo  string    `json:"tributacao_longo_prazo"`
	ValorPatriminioLiq    float64   `json:"valor_patriminio_liquido"`
}
