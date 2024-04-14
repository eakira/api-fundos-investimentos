package domain

type LaminaDomain struct {
	AnoAnterExemplo           string
	AnoExemplo                string
	AnoSemRentab              string
	CalcRentabFundo           string
	CalcRentabFundoGatilho    string
	ClasseRiscoAdmin          string
	FundoCnpj                 string
	CondicCaren               string
	CondicEntr                string
	CondicSaida               string
	ConflitoVenda             string
	ConversaoCotaCanc         string
	ConversaoCotaCompra       string
	DenomSocial               string
	DerivProtecaoCarteira     string
	DistribGestorUnico        string
	DsRentabGatilho           string
	DtComptc                  string
	DtFimDespesa              string
	DtIniAtiv5Ano             string
	DtIniDespesa              string
	EnderEletronico           string
	EnderEletronicoDespesa    string
	EnderEletronicoReclamacao string
	HoraAplicResgate          string
	IndiceRefer               string
	InvestAdic                string
	InvestInicialMin          string
	NmFantasia                string
	Objetivo                  string
	PolitInvest               string
	PrAtivoEmissor            string
	PrPlAlavanc               string
	PrPlAplicMaxFundoUnico    string
	PrPlAtivoCredPriv         string
	PrPlAtivoExterior         string
	PrPlDespesa               string
	PrRentabFundo5Ano         string
	PrVariacaoIndiceRefer5Ano string
	PrVariacaoPerfm           string
	PublicoAlvo               string
	QtAnoPerda                string
	QtDiaCaren                string
	QtDiaConversaoCotaCompra  string
	QtDiaConversaoCotaResgate string
	QtDiaPagtoResgate         string
	QtDiaSaida                string
	RemunDistrib              string
	RentabGatilho             string
	ResgateMin                string
	RestrInvest               string
	RiscoPerda                string
	RiscoPerdaNegativo        string
	TaxaAdm                   string
	TaxaAdmMax                string
	TaxaAdmMin                string
	TaxaAdmObs                string
	TaxaEntr                  string
	TaxaPerfm                 string
	TaxaSaida                 string
	TelSac                    string
	TpDiaPagtoResgate         string
	TpTaxaAdm                 string
	VlAjustePerfmExemplo      string
	VlDespesa3Ano             string
	VlDespesa5Ano             string
	VlDespesaExemplo          string
	VlImpostoExemplo          string
	VlMinPerman               string
	VlPatrimLiq               string
	VlResgateExemplo          string
	VlRetorno3Ano             string
	VlRetorno5Ano             string
	VlTaxaEntrExemplo         string
	VlTaxaSaidaExemplo        string
	InfSac                    string
}

type LaminaCarteiraDomain struct {
	FundoCnpj   string
	DenomSocial string
	DtComptc    string
	PrPlAtivo   string
	TpAtivo     string
}

type LaminaRentabilidadeAnoDomain struct {
	AnoReferencia            string
	FundoCnpj                string
	DenomSocial              string
	DtComptc                 string
	PrPerfIndiceReferAno     string
	PrRentabAno              string
	PrVariacaoIndiceReferAno string
	RentabAnoObservacoes     string
}

type LaminaRentabilidadeMesDomain struct {
	FundoCnpj                string
	DenomSocial              string
	DtComptc                 string
	MesReferencia            string
	PrPerfIndiceReferMes     string
	PrRentabMes              string
	PrVariacaoIndiceReferMes string
	RentabMesObservacoes     string
}
