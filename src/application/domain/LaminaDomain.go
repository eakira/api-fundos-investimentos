package domain

import "time"

type LaminaDomain struct {
	AnoAnterExemplo           int
	AnoExemplo                int
	AnoSemRentab              string
	CalcRentabFundo           string
	CalcRentabFundoGatilho    string
	ClasseRiscoAdmin          int
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
	DtComptc                  time.Time
	DtFimDespesa              time.Time
	DtIniAtiv5Ano             time.Time
	DtIniDespesa              time.Time
	EnderEletronico           string
	EnderEletronicoDespesa    string
	EnderEletronicoReclamacao string
	HoraAplicResgate          string
	IndiceRefer               string
	InvestAdic                float64
	InvestInicialMin          float64
	NmFantasia                string
	Objetivo                  string
	PolitInvest               string
	PrAtivoEmissor            float64
	PrPlAlavanc               float64
	PrPlAplicMaxFundoUnico    float64
	PrPlAtivoCredPriv         float64
	PrPlAtivoExterior         float64
	PrPlDespesa               float64
	PrRentabFundo5Ano         float64
	PrVariacaoIndiceRefer5Ano float64
	PrVariacaoPerfm           float64
	PublicoAlvo               string
	QtAnoPerda                int
	QtDiaCaren                int
	QtDiaConversaoCotaCompra  int
	QtDiaConversaoCotaResgate int
	QtDiaPagtoResgate         int
	QtDiaSaida                int
	RemunDistrib              string
	RentabGatilho             string
	ResgateMin                float64
	RestrInvest               string
	RiscoPerda                string
	RiscoPerdaNegativo        string
	TaxaAdm                   float64
	TaxaAdmMax                float64
	TaxaAdmMin                float64
	TaxaAdmObs                string
	TaxaEntr                  float64
	TaxaPerfm                 float64
	TaxaSaida                 float64
	TelSac                    string
	TpDiaPagtoResgate         string
	TpTaxaAdm                 string
	VlAjustePerfmExemplo      float64
	VlDespesa3Ano             float64
	VlDespesa5Ano             float64
	VlDespesaExemplo          float64
	VlImpostoExemplo          float64
	VlMinPerman               float64
	VlPatrimLiq               float64
	VlResgateExemplo          float64
	VlRetorno3Ano             float64
	VlRetorno5Ano             float64
	VlTaxaEntrExemplo         float64
	VlTaxaSaidaExemplo        float64
	InfSac                    string
}

type LaminaCarteiraDomain struct {
	FundoCnpj   string
	DenomSocial string
	DtComptc    time.Time
	PrPlAtivo   float64
	TpAtivo     string
}

type LaminaRentabilidadeAnoDomain struct {
	AnoReferencia            int
	FundoCnpj                string
	DenomSocial              string
	DtComptc                 time.Time
	PrPerfIndiceReferAno     float64
	PrRentabAno              float64
	PrVariacaoIndiceReferAno float64
	RentabAnoObservacoes     string
}

type LaminaRentabilidadeMesDomain struct {
	FundoCnpj                string
	DenomSocial              string
	DtComptc                 time.Time
	MesReferencia            int
	PrPerfIndiceReferMes     float64
	PrRentabMes              float64
	PrVariacaoIndiceReferMes float64
	RentabMesObservacoes     string
}
