package domain

import "time"

type InformacoesFundoDomain struct {
	AgenciaRisco               string
	ApresentacaoAdministrador  string
	ApresentacaoGestor         string
	CnpjAgenciaRisco           string
	FundoCnpj                  string
	DisclaimerAgenciaRisco     string
	DistribuidorLigado         string
	DataCompetencia            time.Time
	GrauRisco                  string
	InformacoesAutorregulacao  string
	OutrasInformacoes          string
	PeriodoMinimoDivulgacao    string
	PoliticaAdministracaoRisco string
	PoliticaDistribuicaoCotas  string
	PoliticaVoto               string
	RiscoCarteira              string
	ExistenciaAgenciaRisco     string
	TipoFundo                  string
	TributacaoFundoCotista     string
	VotoGestor                 string
}

type InformacoesDivulgacaoDomain struct {
	FundoCnpj        string
	DataCompetencia  time.Time
	FormaInformacoes string
	LocalInformacoes string
	MeioInformacoes  string
	TipoFundo        string
}

type InformacoesCotistaDomain struct {
	FundoCnpj            string
	DescricaoRespCotista string
	DataCompetencia      time.Time
	FormaInfCotista      string
	LocalInfCotista      string
	MeioInfCotista       string
	TipoFundo            string
}

type ServicoPrestadoDomain struct {
	FundoCnpj        string
	DescricaoServico string
	DataCompetencia  time.Time
	NomePrestador    string
	TipoFundo        string
}
