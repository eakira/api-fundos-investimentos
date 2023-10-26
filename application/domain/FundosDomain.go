package domain

import "time"

type FundosDomain struct {
	Id                    string
	AdministradorNome     string
	AuditorNome           string
	CodigoCVM             int16
	Classe                string
	ClasseAnbima          string
	AdministradorCNPJ     string
	AuditorCNPJ           string
	ControladorCNPJ       string
	CustodianteCNPJ       string
	FundoCNPJ             string
	Condominio            string
	Controlador           string
	GestorCPFCNPJ         string
	Custodiante           string
	FundoNome             string
	Diretor               string
	DataCancelamento      time.Time
	DataConstituicao      time.Time
	DataFim               time.Time
	DataInicio            time.Time
	DataInicioClasse      time.Time
	DataInicioSocial      time.Time
	DataInicioSituacao    time.Time
	DataPatrimonioLiquido time.Time
	DataRegistro          time.Time
	Entidade              string
	CotasPossui           string
	Exclusivo             string
	GestorNome            string
	InformacoesAdicionais string
	InfoTaxaPerformance   string
	Exterior100           string
	PessoaFouPJ           string
	PublicoAlvo           string
	IndicadorDesempenho   string
	Situacao              string
	TaxaAdm               float64
	TaxaPerformance       float64
	TipoFundo             string
	TributacaoLongoPrazo  string
	ValorPatriminioLiq    float64
}
