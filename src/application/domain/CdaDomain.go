package domain

import "time"

type CdaBlc1Domain struct {
	CodigoConta              string
	FundoCnpj                string
	DataCompetencia          string
	PlanoContabil            string
	SaldoConta               string
	CdIsin                   string
	CdSelic                  string
	DenominacaoSocial        string
	DataEmissao              string
	DataVencimento           string
	DataCompetenciaDoc       string
	PrazoConfidencialidade   string
	EmissorLigado            string
	QuantidadeAquisicoes     string
	QuantidadePosicaoFinal   string
	QuantidadeVendas         string
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	TipoTituloPublico        string
	ValorAquisicoes          string
	ValorCustoPosicaoFinal   string
	ValorMercadoPosicaoFinal string
	ValorVendas              string
}

type CdaBlc2Domain struct {
	CnpjFundo                string
	CnpjFundoCota            string
	DenominacaoSocial        string
	DataCompetencia          string
	PrazoConfidencialidade   string
	EmissorLigado            string
	NomeFundoCota            string
	QuantidadeAquisicoes     string
	QuantidadePosicaoFinal   string
	QuantidadeVendas         string
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	ValorAquisicoes          string
	ValorCustoPosicaoFinal   string
	ValorMercadoPosicaoFinal string
	ValorVendas              string
}

type CdaBlc3Domain struct {
	CodigoSwap               string
	CnpjFundo                string
	DenominacaoSocial        string
	DescricaoTipoAtivo       string
	DataCompetencia          string
	PrazoConfidencialidade   string
	EmissorLigado            string
	QuantidadeAquisicoes     string
	QuantidadePosicaoFinal   string
	QuantidadeVendas         string
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	ValorAquisicoes          string
	ValorCustoPosicaoFinal   string
	ValorMercadoPosicaoFinal string
	ValorVendas              string
}

type CdaBlc4Domain struct {
	CodigoAtivo              string
	CodigoISIN               string
	CnpjFundo                string
	DenominacaoSocial        string
	DescricaoAtivo           string
	DataCompetencia          string
	PrazoConfidencialidade   time.Time
	DataFimVigencia          time.Time
	DataInicioVigencia       time.Time
	EmissorLigado            string
	QuantidadeAquisicoes     string
	QuantidadePosicaoFinal   string
	QuantidadeVendas         string
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	ValorAquisicoes          string
	ValorCustoPosicaoFinal   string
	ValorMercadoPosicaoFinal string
	ValorVendas              string
}

type CdaBlc5Domain struct {
	NomeAgenciaClassificacaoRisco string
	CodigoIndexadorPosFixados     string
	CnpjEmissor                   string
	CnpjFundo                     string
	DenominacaoSocial             string
	DescricaoIndexadorPosFixados  string
	DataCompetencia               string
	PrazoConfidencialidade        time.Time
	DataClassificacaoRisco        time.Time
	DataVencimento                time.Time
	NomeEmissor                   string
	EmissorLigado                 string
	GrauRisco                     string
	PercentualCupomPosFixados     float64
	PercentualIndexadorPosFixados float64
	PercentualTaxaPrefixada       float64
	QuantidadeAquisicoes          float64
	QuantidadePosicaoFinal        float64
	QuantidadeVendas              float64
	RiscoEmissor                  string
	TituloPosFixado               string
	TipoAplicacao                 string
	TipoAtivo                     string
	TipoFundo                     string
	TipoNegociacao                string
	ValorAquisicoes               float64
	ValorCustoPosicaoFinal        float64
	ValorMercadoPosicaoFinal      float64
	ValorVendas                   float64
}

type CdaBlc6Domain struct {
	CodigoIndexadorPosFixados      string
	CnpjFundo                      string
	CnpjInstituicaoFinanceiraCoobr string
	CpfCnpjEmissor                 string
	DenominacaoSocial              string
	DescricaoIndexadorPosFixados   string
	DataCompetencia                string
	PrazoConfidencialidade         time.Time
	DataVencimento                 time.Time
	NomeEmissor                    string
	EmissorLigado                  string
	PessoaFisicaJuridicaEmissor    string
	PercentualCupomPosFixados      float64
	PercentualIndexadorPosFixados  float64
	PercentualTaxaPrefixada        float64
	QuantidadeAquisicoes           float64
	QuantidadePosicaoFinal         float64
	QuantidadeVendas               float64
	TituloRegistradoCetip          string
	TituloPossuiGarantia           string
	TituloPosFixado                string
	TipoAplicacao                  string
	TipoAtivo                      string
	TipoFundo                      string
	TipoNegociacao                 string
	ValorAquisicoes                float64
	ValorCustoPosicaoFinal         float64
	ValorMercadoPosicaoFinal       float64
	ValorVendas                    float64
}

type CdaBlc7Domain struct {
	AgenciaRisco                      string
	BolsaMercado                      string
	CodigoAtivoBolsaMercado           string
	CodigoBolsaMercado                string
	CodigoPais                        string
	CnpjFundo                         string
	DenominacaoSocial                 string
	DescricaoAtivoExterior            string
	DataCompetencia                   time.Time
	PrazoConfidencialidade            time.Time
	DataClassificacaoRisco            time.Time
	DataVencimento                    time.Time
	NomeEmissor                       string
	EmissorLigado                     string
	GrauRisco                         string
	InvestimentoColetivo              string
	GestaoVeiculoInvestimentoColetivo string
	Pais                              string
	QuantidadeAquisicoes              float64
	QuantidadeAtivosExterior          int
	QuantidadePosicaoFinal            float64
	QuantidadeVendas                  float64
	RiscoEmissor                      string
	TipoAplicacao                     string
	TipoAtivo                         string
	TipoFundo                         string
	TipoNegociacao                    string
	ValorAquisicoes                   float64
	ValorAtivoExterior                float64
	ValorCustoPosicaoFinal            float64
	ValorMercadoPosicaoFinal          float64
	ValorVendas                       float64
}

type CdaBlc8Domain struct {
	CnpjFundo                   string
	CpfCnpjEmissor              string
	DenominacaoSocial           string
	DescricaoAtivo              string
	DataCompetencia             time.Time
	PrazoConfidencialidade      time.Time
	NomeEmissor                 string
	EmissorLigado               string
	PessoaFisicaJuridicaEmissor string
	QuantidadeAquisicoes        float64
	QuantidadePosicaoFinal      float64
	QuantidadeVendas            float64
	TipoAplicacao               string
	TipoAtivo                   string
	TipoFundo                   string
	TipoNegociacao              string
	ValorAquisicoes             float64
	ValorCustoPosicaoFinal      float64
	ValorMercadoPosicaoFinal    float64
	ValorVendas                 float64
}

type CdaConfidencialDomain struct {
	CnpjFundo           string
	DenominacaoSocial   string
	DataCompetencia     time.Time
	PrazoConfidencial   time.Time
	TipoAplicacao       string
	TipoFundo           string
	ValorAquisicoes     float64
	ValorCustoPosicao   float64
	ValorMercadoPosicao float64
	ValorVendas         float64
}

type CdaFiimDomain struct {
	BolsaMercado             string
	CodigoAtivo              string
	CodigoBolsaMercado       string
	CodigoPais               int
	CodigoSELIC              string
	CnpjFundo                string
	CpfCnpjEmissor           int
	DenominacaoSocial        string
	DescricaoAtivo           string
	DataCompetencia          time.Time
	PrazoConfidencial        time.Time
	DataInicioVigencia       time.Time
	DataVencimento           time.Time
	Emissor                  string
	EmissorLigado            string
	IDDocumento              int
	Pais                     string
	PessoaFisicaJuridica     string
	QuantidadeAquisicoes     float64
	QuantidadePosicaoFinal   float64
	QuantidadeVendas         float64
	RiscoEmissor             string
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	ValorAquisicoes          float64
	ValorCustoPosicaoFinal   float64
	ValorMercadoPosicaoFinal float64
	ValorPatrimonioLiquido   float64
	ValorVendas              float64
}

type CdaFiimConfidencialDomain struct {
	CnpjFundo                string
	DenominacaoSocial        string
	DataCompetencia          time.Time
	PrazoConfidencial        time.Time
	IDDocumento              int
	TipoAplicacao            string
	TipoFundo                string
	ValorAquisicoes          float64
	ValorCustoPosicaoFinal   float64
	ValorMercadoPosicaoFinal float64
	ValorVendas              float64
}

type CdaPatrimonioLiquidoDomain struct {
	CnpjFundo         string
	DenominacaoSocial string
	DataCompetencia   time.Time
	TipoFundo         string
	ValorPatrimonio   float64
}
