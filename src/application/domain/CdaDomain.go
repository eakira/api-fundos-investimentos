package domain

type CdaSelicDomain struct {
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
	PrazoConfidencialidade   string
	DataFimVigencia          string
	DataInicioVigencia       string
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
	PrazoConfidencialidade        string
	DataClassificacaoRisco        string
	DataVencimento                string
	NomeEmissor                   string
	EmissorLigado                 string
	GrauRisco                     string
	PercentualCupomPosFixados     string
	PercentualIndexadorPosFixados string
	PercentualTaxaPrefixada       string
	QuantidadeAquisicoes          string
	QuantidadePosicaoFinal        string
	QuantidadeVendas              string
	RiscoEmissor                  string
	TituloPosFixado               string
	TipoAplicacao                 string
	TipoAtivo                     string
	TipoFundo                     string
	TipoNegociacao                string
	ValorAquisicoes               string
	ValorCustoPosicaoFinal        string
	ValorMercadoPosicaoFinal      string
	ValorVendas                   string
}

type CdaBlc6Domain struct {
	CodigoIndexadorPosFixados      string
	CnpjFundo                      string
	CnpjInstituicaoFinanceiraCoobr string
	CpfCnpjEmissor                 string
	DenominacaoSocial              string
	DescricaoIndexadorPosFixados   string
	DataCompetencia                string
	PrazoConfidencialidade         string
	DataVencimento                 string
	NomeEmissor                    string
	EmissorLigado                  string
	PessoaFisicaJuridicaEmissor    string
	PercentualCupomPosFixados      string
	PercentualIndexadorPosFixados  string
	PercentualTaxaPrefixada        string
	QuantidadeAquisicoes           string
	QuantidadePosicaoFinal         string
	QuantidadeVendas               string
	TituloRegistradoCetip          string
	TituloPossuiGarantia           string
	TituloPosFixado                string
	TipoAplicacao                  string
	TipoAtivo                      string
	TipoFundo                      string
	TipoNegociacao                 string
	ValorAquisicoes                string
	ValorCustoPosicaoFinal         string
	ValorMercadoPosicaoFinal       string
	ValorVendas                    string
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
	DataCompetencia                   string
	PrazoConfidencialidade            string
	DataClassificacaoRisco            string
	DataVencimento                    string
	NomeEmissor                       string
	EmissorLigado                     string
	GrauRisco                         string
	InvestimentoColetivo              string
	GestaoVeiculoInvestimentoColetivo string
	Pais                              string
	QuantidadeAquisicoes              string
	QuantidadeAtivosExterior          string
	QuantidadePosicaoFinal            string
	QuantidadeVendas                  string
	RiscoEmissor                      string
	TipoAplicacao                     string
	TipoAtivo                         string
	TipoFundo                         string
	TipoNegociacao                    string
	ValorAquisicoes                   string
	ValorAtivoExterior                string
	ValorCustoPosicaoFinal            string
	ValorMercadoPosicaoFinal          string
	ValorVendas                       string
}

type CdaBlc8Domain struct {
	CnpjFundo                   string
	CpfCnpjEmissor              string
	DenominacaoSocial           string
	DescricaoAtivo              string
	DataCompetencia             string
	PrazoConfidencialidade      string
	NomeEmissor                 string
	EmissorLigado               string
	PessoaFisicaJuridicaEmissor string
	QuantidadeAquisicoes        string
	QuantidadePosicaoFinal      string
	QuantidadeVendas            string
	TipoAplicacao               string
	TipoAtivo                   string
	TipoFundo                   string
	TipoNegociacao              string
	ValorAquisicoes             string
	ValorCustoPosicaoFinal      string
	ValorMercadoPosicaoFinal    string
	ValorVendas                 string
}

type CdaConfidencialDomain struct {
	CnpjFundo           string
	DenominacaoSocial   string
	DataCompetencia     string
	PrazoConfidencial   string
	TipoAplicacao       string
	TipoFundo           string
	ValorAquisicoes     string
	ValorCustoPosicao   string
	ValorMercadoPosicao string
	ValorVendas         string
}

type CdaFiimDomain struct {
	BolsaMercado             string
	CodigoAtivo              string
	CodigoBolsaMercado       string
	CodigoPais               string
	CodigoSELIC              string
	CnpjFundo                string
	CpfCnpjEmissor           string
	DenominacaoSocial        string
	DescricaoAtivo           string
	DataCompetencia          string
	PrazoConfidencial        string
	DataInicioVigencia       string
	DataVencimento           string
	Emissor                  string
	EmissorLigado            string
	IDDocumento              string
	Pais                     string
	PessoaFisicaJuridica     string
	QuantidadeAquisicoes     string
	QuantidadePosicaoFinal   string
	QuantidadeVendas         string
	RiscoEmissor             string
	TipoAplicacao            string
	TipoAtivo                string
	TipoFundo                string
	TipoNegociacao           string
	ValorAquisicoes          string
	ValorCustoPosicaoFinal   string
	ValorMercadoPosicaoFinal string
	ValorPatrimonioLiquido   string
	ValorVendas              string
}

type CdaFiimConfidencialDomain struct {
	CnpjFundo                string
	DenominacaoSocial        string
	DataCompetencia          string
	PrazoConfidencial        string
	IDDocumento              string
	TipoAplicacao            string
	TipoFundo                string
	ValorAquisicoes          string
	ValorCustoPosicaoFinal   string
	ValorMercadoPosicaoFinal string
	ValorVendas              string
}

type CdaPatrimonioLiquidoDomain struct {
	CnpjFundo         string
	DenominacaoSocial string
	DataCompetencia   string
	TipoFundo         string
	ValorPatrimonio   string
}
