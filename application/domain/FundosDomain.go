package domain

import "time"

type FundosDomain struct {
	administradorNome     string
	auditorNome           string
	codigoCVM             int16
	classe                string
	classeAnbima          string
	administradorCNPJ     string
	auditorCNPJ           string
	controladorCNPJ       string
	custodianteCNPJ       string
	fundoCNPJ             string
	condominio            string
	controlador           string
	gestorCPFCNPJ         string
	custodiante           string
	fundoNome             string
	diretor               string
	dataCancelamento      time.Time
	dataConstituicao      time.Time
	dataFim               time.Time
	dataInicio            time.Time
	dataInicioClasse      time.Time
	dataInicioSocial      time.Time
	dataInicioSituacao    time.Time
	dataPatrimonioLiquido time.Time
	dataRegistro          time.Time
	entidade              string
	cotasPossui           string
	exclusivo             string
	gestorNome            string
	informacoesAdicionais string
	InfoTaxaPerformance   string
	exterior100           string
	PessoaFOuPJ           string
	publicoAlvo           string
	indicadorDesempenho   string
	situacao              string
	taxaAdm               float64
	taxaPerformance       float64
	tipoFundo             string
	tributacaoLongoPrazo  string
	valorPatriminioLiq    float64
}
