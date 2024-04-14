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

type CdaFundosInvestimentosDomain struct {
	FundoCnpj                string
	FundoCnpjCota            string
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

type CdaSwapDomain struct {
	CodigoSwap               string
	FundoCnpj                string
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

type CdaDemaisAtivosDomain struct {
	CodigoAtivo              string
	CodigoISIN               string
	FundoCnpj                string
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

type CdaDepositoAPrazoOutrosAtivosDomain struct {
	NomeAgenciaClassificacaoRisco string
	CodigoIndexadorPosFixados     string
	CnpjEmissor                   string
	FundoCnpj                     string
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

type CdaAgroCreditoPrivadoDomain struct {
	CodigoIndexadorPosFixados      string
	FundoCnpj                      string
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

type CdaInvestimentosExteriorDomain struct {
	AgenciaRisco                      string
	BolsaMercado                      string
	CodigoAtivoBolsaMercado           string
	CodigoBolsaMercado                string
	CodigoPais                        string
	FundoCnpj                         string
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

type CdaDemaisAtivosNaoCodificadosDomain struct {
	FundoCnpj                   string
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
	FundoCnpj           string
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
	FundoCnpj                string
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
	FundoCnpj                string
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
	FundoCnpj         string
	DenominacaoSocial string
	DataCompetencia   string
	TipoFundo         string
	ValorPatrimonio   string
}
