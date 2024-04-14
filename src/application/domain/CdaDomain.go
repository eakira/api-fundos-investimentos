package domain

import "time"

type CdaSelicDomain struct {
	CodigoConta              string
	FundoCnpj                string
	DataCompetencia          time.Time
	PlanoContabil            string
	SaldoConta               string
	CdIsin                   string
	CdSelic                  string
	DenominacaoSocial        string
	DataEmissao              time.Time
	DataVencimento           time.Time
	DataCompetenciaDoc       time.Time
	PrazoConfidencialidade   time.Time
	EmissorLigado            string
	QuantidadeAquisicoes     float64
	QuantidadePosicaoFinal   float64
	QuantidadeVendas         float64
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	TipoTituloPublico        string
	ValorAquisicoes          float64
	ValorCustoPosicaoFinal   float64
	ValorMercadoPosicaoFinal float64
	ValorVendas              float64
}

type CdaFundosInvestimentosDomain struct {
	FundoCnpj                string
	FundoCnpjCota            string
	DenominacaoSocial        string
	DataCompetencia          time.Time
	PrazoConfidencialidade   time.Time
	EmissorLigado            string
	NomeFundoCota            string
	QuantidadeAquisicoes     float64
	QuantidadePosicaoFinal   float64
	QuantidadeVendas         float64
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	ValorAquisicoes          float64
	ValorCustoPosicaoFinal   float64
	ValorMercadoPosicaoFinal float64
	ValorVendas              float64
}

type CdaSwapDomain struct {
	CodigoSwap               string
	FundoCnpj                string
	DenominacaoSocial        string
	DescricaoTipoAtivo       string
	DataCompetencia          time.Time
	PrazoConfidencialidade   time.Time
	EmissorLigado            string
	QuantidadeAquisicoes     float64
	QuantidadePosicaoFinal   float64
	QuantidadeVendas         float64
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	ValorAquisicoes          float64
	ValorCustoPosicaoFinal   float64
	ValorMercadoPosicaoFinal float64
	ValorVendas              float64
}

type CdaDemaisAtivosDomain struct {
	CodigoAtivo              string
	CodigoISIN               string
	FundoCnpj                string
	DenominacaoSocial        string
	DescricaoAtivo           string
	DataCompetencia          time.Time
	PrazoConfidencialidade   time.Time
	DataFimVigencia          time.Time
	DataInicioVigencia       time.Time
	EmissorLigado            string
	QuantidadeAquisicoes     float64
	QuantidadePosicaoFinal   float64
	QuantidadeVendas         float64
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	ValorAquisicoes          float64
	ValorCustoPosicaoFinal   float64
	ValorMercadoPosicaoFinal float64
	ValorVendas              float64
}

type CdaDepositoAPrazoOutrosAtivosDomain struct {
	NomeAgenciaClassificacaoRisco string
	CodigoIndexadorPosFixados     string
	CnpjEmissor                   string
	FundoCnpj                     string
	DenominacaoSocial             string
	DescricaoIndexadorPosFixados  string
	DataCompetencia               time.Time
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

type CdaAgroCreditoPrivadoDomain struct {
	CodigoIndexadorPosFixados      string
	FundoCnpj                      string
	CnpjInstituicaoFinanceiraCoobr string
	CpfCnpjEmissor                 string
	DenominacaoSocial              string
	DescricaoIndexadorPosFixados   string
	DataCompetencia                time.Time
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

type CdaInvestimentosExteriorDomain struct {
	AgenciaRisco                      string
	BolsaMercado                      string
	CodigoAtivoBolsaMercado           string
	CodigoBolsaMercado                string
	CodigoPais                        string
	FundoCnpj                         string
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
	QuantidadeAtivosExterior          float64
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

type CdaDemaisAtivosNaoCodificadosDomain struct {
	FundoCnpj                   string
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
	FundoCnpj           string
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
	FundoCnpj                string
	CpfCnpjEmissor           string
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
	FundoCnpj                string
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
	FundoCnpj         string
	DenominacaoSocial string
	DataCompetencia   time.Time
	TipoFundo         string
	ValorPatrimonio   float64
}
