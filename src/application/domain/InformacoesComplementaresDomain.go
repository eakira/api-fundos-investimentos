package domain

type InformacoesFundoDomain struct {
	AgenciaRisco               string
	ApresentacaoAdministrador  string
	ApresentacaoGestor         string
	CnpjAgenciaRisco           string
	CnpjFundo                  string
	DisclaimerAgenciaRisco     string
	DistribuidorLigado         string
	DataCompetencia            string
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
	CnpjFundo        string
	DataCompetencia  string
	FormaInformacoes string
	LocalInformacoes string
	MeioInformacoes  string
	TipoFundo        string
}

type InformacoesCotistaDomain struct {
	CnpjFundo            string
	DescricaoRespCotista string
	DataCompetencia      string
	FormaInfCotista      string
	LocalInfCotista      string
	MeioInfCotista       string
	TipoFundo            string
}

type ServicoPrestadoDomain struct {
	CnpjFundo        string
	DescricaoServico string
	DataCompetencia  string
	NomePrestador    string
	TipoFundo        string
}
