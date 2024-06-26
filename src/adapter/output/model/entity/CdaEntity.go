package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CdaSelicEntity struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	CodigoConta              string             `bson:"codigo_conta"`
	FundoCnpj                string             `bson:"fundo_cnpj"`
	DataCompetencia          time.Time          `bson:"data_competencia"`
	PlanoContabil            string             `bson:"plano_contabil"`
	SaldoConta               string             `bson:"saldo_conta"`
	CdIsin                   string             `bson:"cd_isin"`
	CdSelic                  string             `bson:"cd_selic"`
	DenominacaoSocial        string             `bson:"denominacao_social"`
	DataEmissao              time.Time          `bson:"data_emissao"`
	DataVencimento           time.Time          `bson:"data_vencimento"`
	DataCompetenciaDoc       time.Time          `bson:"data_competencia_doc"`
	PrazoConfidencialidade   time.Time          `bson:"prazo_confidencialidade"`
	EmissorLigado            string             `bson:"emissor_ligado"`
	QuantidadeAquisicoes     float64            `bson:"quantidade_aquisicoes"`
	QuantidadePosicaoFinal   float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas         float64            `bson:"quantidade_vendas"`
	TipoAplicacao            string             `bson:"tipo_aplicacao"`
	TipoAtivo                string             `bson:"tipo_ativo"`
	TipoFundo                string             `bson:"tipo_fundo"`
	TipoNegociacao           string             `bson:"tipo_negociacao"`
	TipoTituloPublico        string             `bson:"tipo_titulo_publico"`
	ValorAquisicoes          float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal   float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal float64            `bson:"valor_mercado_posicao_final"`
	ValorVendas              float64            `bson:"valor_vendas"`
}

type CdaFundosInvestimentosEntity struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	FundoCnpj                string             `bson:"cnpj_fundo"`
	FundoCnpjCota            string             `bson:"cnpj_fundo_cota"`
	DenominacaoSocial        string             `bson:"denominacao_social"`
	DataCompetencia          time.Time          `bson:"data_competencia"`
	PrazoConfidencialidade   time.Time          `bson:"prazo_confidencialidade"`
	EmissorLigado            string             `bson:"emissor_ligado"`
	NomeFundoCota            string             `bson:"nome_fundo_cota"`
	QuantidadeAquisicoes     float64            `bson:"quantidade_aquisicoes"`
	QuantidadePosicaoFinal   float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas         float64            `bson:"quantidade_vendas"`
	TipoAplicacao            string             `bson:"tipo_aplicacao"`
	TipoAtivo                string             `bson:"tipo_ativo"`
	TipoFundo                string             `bson:"tipo_fundo"`
	TipoNegociacao           string             `bson:"tipo_negociacao"`
	ValorAquisicoes          float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal   float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal float64            `bson:"valor_mercado_posicao_final"`
	ValorVendas              float64            `bson:"valor_vendas"`
}

type CdaSwapEntity struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	CodigoSwap               string             `bson:"codigo_swap"`
	FundoCnpj                string             `bson:"cnpj_fundo"`
	DenominacaoSocial        string             `bson:"denominacao_social"`
	DescricaoTipoAtivo       string             `bson:"descricao_tipo_ativo"`
	DataCompetencia          time.Time          `bson:"data_competencia"`
	PrazoConfidencialidade   time.Time          `bson:"prazo_confidencialidade"`
	EmissorLigado            string             `bson:"emissor_ligado"`
	QuantidadeAquisicoes     float64            `bson:"quantidade_aquisicoes"`
	QuantidadePosicaoFinal   float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas         float64            `bson:"quantidade_vendas"`
	TipoAplicacao            string             `bson:"tipo_aplicacao"`
	TipoAtivo                string             `bson:"tipo_ativo"`
	TipoFundo                string             `bson:"tipo_fundo"`
	TipoNegociacao           string             `bson:"tipo_negociacao"`
	ValorAquisicoes          float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal   float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal float64            `bson:"valor_mercado_posicao_final"`
	ValorVendas              float64            `bson:"valor_vendas"`
}

type CdaDemaisAtivosEntity struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	CodigoAtivo              string             `bson:"codigo_ativo"`
	CodigoISIN               string             `bson:"codigo_isin"`
	FundoCnpj                string             `bson:"cnpj_fundo"`
	DenominacaoSocial        string             `bson:"denominacao_social"`
	DescricaoAtivo           string             `bson:"descricao_ativo"`
	DataCompetencia          time.Time          `bson:"data_competencia"`
	PrazoConfidencialidade   time.Time          `bson:"prazo_confidencialidade"`
	DataFimVigencia          time.Time          `bson:"data_fim_vigencia"`
	DataInicioVigencia       time.Time          `bson:"data_inicio_vigencia"`
	EmissorLigado            string             `bson:"emissor_ligado"`
	QuantidadeAquisicoes     float64            `bson:"quantidade_aquisicoes"`
	QuantidadePosicaoFinal   float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas         float64            `bson:"quantidade_vendas"`
	TipoAplicacao            string             `bson:"tipo_aplicacao"`
	TipoAtivo                string             `bson:"tipo_ativo"`
	TipoFundo                string             `bson:"tipo_fundo"`
	TipoNegociacao           string             `bson:"tipo_negociacao"`
	ValorAquisicoes          float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal   float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal float64            `bson:"valor_mercado_posicao_final"`
	ValorVendas              float64            `bson:"valor_vendas"`
}

type CdaDepositoAPrazoOutrosAtivosEntity struct {
	ID                            primitive.ObjectID `bson:"_id,omitempty"`
	NomeAgenciaClassificacaoRisco string             `bson:"nome_agencia_classificacao_risco"`
	CodigoIndexadorPosFixados     string             `bson:"codigo_indexador_pos_fixados"`
	CnpjEmissor                   string             `bson:"cnpj_emissor"`
	FundoCnpj                     string             `bson:"cnpj_fundo"`
	DenominacaoSocial             string             `bson:"denominacao_social"`
	DescricaoIndexadorPosFixados  string             `bson:"descricao_indexador_pos_fixados"`
	DataCompetencia               time.Time          `bson:"data_competencia"`
	PrazoConfidencialidade        time.Time          `bson:"prazo_confidencialidade"`
	DataClassificacaoRisco        time.Time          `bson:"data_classificacao_risco"`
	DataVencimento                time.Time          `bson:"data_vencimento"`
	NomeEmissor                   string             `bson:"nome_emissor"`
	EmissorLigado                 string             `bson:"emissor_ligado"`
	GrauRisco                     string             `bson:"grau_risco"`
	PercentualCupomPosFixados     float64            `bson:"percentual_cupom_pos_fixados"`
	PercentualIndexadorPosFixados float64            `bson:"percentual_indexador_pos_fixados"`
	PercentualTaxaPrefixada       float64            `bson:"percentual_taxa_prefixada"`
	QuantidadeAquisicoes          float64            `bson:"quantidade_aquisicoes"`
	QuantidadePosicaoFinal        float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas              float64            `bson:"quantidade_vendas"`
	RiscoEmissor                  string             `bson:"risco_emissor"`
	TituloPosFixado               string             `bson:"titulo_pos_fixado"`
	TipoAplicacao                 string             `bson:"tipo_aplicacao"`
	TipoAtivo                     string             `bson:"tipo_ativo"`
	TipoFundo                     string             `bson:"tipo_fundo"`
	TipoNegociacao                float64            `bson:"tipo_negociacao"`
	ValorAquisicoes               float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal        float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal      float64            `bson:"valor_mercado_posicao_final"`
}

type CdaAgroCreditoPrivadoEntity struct {
	ID                             primitive.ObjectID `bson:"_id,omitempty"`
	CodigoIndexadorPosFixados      string             `bson:"codigo_indexador_pos_fixados"`
	FundoCnpj                      string             `bson:"cnpj_fundo"`
	CnpjInstituicaoFinanceiraCoobr string             `bson:"cnpj_instituicao_financeira_coobr"`
	CpfCnpjEmissor                 string             `bson:"cpf_cnpj_emissor"`
	DenominacaoSocial              string             `bson:"denominacao_social"`
	DescricaoIndexadorPosFixados   string             `bson:"descricao_indexador_pos_fixados"`
	DataCompetencia                time.Time          `bson:"data_competencia"`
	PrazoConfidencialidade         time.Time          `bson:"prazo_confidencialidade"`
	DataVencimento                 time.Time          `bson:"data_vencimento"`
	NomeEmissor                    string             `bson:"nome_emissor"`
	EmissorLigado                  string             `bson:"emissor_ligado"`
	PessoaFisicaJuridicaEmissor    string             `bson:"pessoa_fisica_juridica_emissor"`
	PercentualCupomPosFixados      float64            `bson:"percentual_cupom_pos_fixados"`
	PercentualIndexadorPosFixados  float64            `bson:"percentual_indexador_pos_fixados"`
	PercentualTaxaPrefixada        float64            `bson:"percentual_taxa_prefixada"`
	QuantidadeAquisicoes           float64            `bson:"quantidade_aquisicoes"`
	QuantidadePosicaoFinal         float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas               float64            `bson:"quantidade_vendas"`
	TituloRegistradoCetip          string             `bson:"titulo_registrado_cetip"`
	TituloPossuiGarantia           string             `bson:"titulo_possui_garantia"`
	TituloPosFixado                string             `bson:"titulo_pos_fixado"`
	TipoAplicacao                  string             `bson:"tipo_aplicacao"`
	TipoAtivo                      string             `bson:"tipo_ativo"`
	TipoFundo                      string             `bson:"tipo_fundo"`
	TipoNegociacao                 string             `bson:"tipo_negociacao"`
	ValorAquisicoes                float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal         float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal       float64            `bson:"valor_mercado_posicao_final"`
	ValorVendas                    float64            `bson:"valor_vendas"`
}

type CdaInvestimentosExteriorEntity struct {
	ID                                primitive.ObjectID `bson:"_id,omitempty"`
	AgenciaRisco                      string             `bson:"agencia_risco"`
	BolsaMercado                      string             `bson:"bolsa_mercado"`
	CodigoAtivoBolsaMercado           string             `bson:"codigo_ativo_bolsa_mercado"`
	CodigoBolsaMercado                string             `bson:"codigo_bolsa_mercado"`
	CodigoPais                        string             `bson:"codigo_pais"`
	FundoCnpj                         string             `bson:"cnpj_fundo"`
	DenominacaoSocial                 string             `bson:"denominacao_social"`
	DescricaoAtivoExterior            string             `bson:"descricao_ativo_exterior"`
	DataCompetencia                   time.Time          `bson:"data_competencia"`
	PrazoConfidencialidade            time.Time          `bson:"prazo_confidencialidade"`
	DataClassificacaoRisco            time.Time          `bson:"data_classificacao_risco"`
	DataVencimento                    time.Time          `bson:"data_vencimento"`
	NomeEmissor                       string             `bson:"nome_emissor"`
	EmissorLigado                     string             `bson:"emissor_ligado"`
	GrauRisco                         string             `bson:"grau_risco"`
	InvestimentoColetivo              string             `bson:"investimento_coletivo"`
	GestaoVeiculoInvestimentoColetivo string             `bson:"gestao_veiculo_investimento_coletivo"`
	Pais                              string             `bson:"pais"`
	QuantidadeAquisicoes              float64            `bson:"quantidade_aquisicoes"`
	QuantidadeAtivosExterior          float64            `bson:"quantidade_ativos_exterior"`
	QuantidadePosicaoFinal            float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas                  float64            `bson:"quantidade_vendas"`
	RiscoEmissor                      string             `bson:"risco_emissor"`
	TipoAplicacao                     string             `bson:"tipo_aplicacao"`
	TipoAtivo                         string             `bson:"tipo_ativo"`
	TipoFundo                         string             `bson:"tipo_fundo"`
	TipoNegociacao                    string             `bson:"tipo_negociacao"`
	ValorAquisicoes                   float64            `bson:"valor_aquisicoes"`
	ValorAtivoExterior                float64            `bson:"valor_ativo_exterior"`
	ValorCustoPosicaoFinal            float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal          float64            `bson:"valor_mercado_posicao_final"`
	ValorVendas                       float64            `bson:"valor_vendas"`
}

type CdaDemaisAtivosNaoCodificadosEntity struct {
	ID                          primitive.ObjectID `bson:"_id,omitempty"`
	FundoCnpj                   string             `bson:"cnpj_fundo"`
	CpfCnpjEmissor              string             `bson:"cpf_cnpj_emissor"`
	DenominacaoSocial           string             `bson:"denominacao_social"`
	DescricaoAtivo              string             `bson:"descricao_ativo"`
	DataCompetencia             time.Time          `bson:"data_competencia"`
	PrazoConfidencialidade      time.Time          `bson:"prazo_confidencialidade"`
	NomeEmissor                 string             `bson:"nome_emissor"`
	EmissorLigado               string             `bson:"emissor_ligado"`
	PessoaFisicaJuridicaEmissor string             `bson:"pessoa_fisica_juridica_emissor"`
	QuantidadeAquisicoes        float64            `bson:"quantidade_aquisicoes"`
	QuantidadePosicaoFinal      float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas            float64            `bson:"quantidade_vendas"`
	TipoAplicacao               string             `bson:"tipo_aplicacao"`
	TipoAtivo                   string             `bson:"tipo_ativo"`
	TipoFundo                   string             `bson:"tipo_fundo"`
	TipoNegociacao              string             `bson:"tipo_negociacao"`
	ValorAquisicoes             float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal      float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal    float64            `bson:"valor_mercado_posicao_final"`
	ValorVendas                 float64            `bson:"valor_vendas"`
}

type CdaConfidencialEntity struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	FundoCnpj           string             `bson:"cnpj_fundo"`
	DenominacaoSocial   string             `bson:"denominacao_social"`
	DataCompetencia     time.Time          `bson:"data_competencia"`
	PrazoConfidencial   time.Time          `bson:"prazo_confidencial"`
	TipoAplicacao       string             `bson:"tipo_aplicacao"`
	TipoFundo           string             `bson:"tipo_fundo"`
	ValorAquisicoes     float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicao   float64            `bson:"valor_custo_posicao"`
	ValorMercadoPosicao float64            `bson:"valor_mercado_posicao"`
	ValorVendas         float64            `bson:"valor_vendas"`
}

type CdaFiimEntity struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	BolsaMercado             string             `bson:"bolsa_mercado"`
	CodigoAtivo              string             `bson:"codigo_ativo"`
	CodigoBolsaMercado       string             `bson:"codigo_bolsa_mercado"`
	CodigoPais               int                `bson:"codigo_pais"`
	CodigoSELIC              string             `bson:"codigo_selic"`
	FundoCnpj                string             `bson:"cnpj_fundo"`
	CpfCnpjEmissor           string             `bson:"cpf_cnpj_emissor"`
	DenominacaoSocial        string             `bson:"denominacao_social"`
	DescricaoAtivo           string             `bson:"descricao_ativo"`
	DataCompetencia          time.Time          `bson:"data_competencia"`
	PrazoConfidencial        time.Time          `bson:"prazo_confidencial"`
	DataInicioVigencia       time.Time          `bson:"data_inicio_vigencia"`
	DataVencimento           time.Time          `bson:"data_vencimento"`
	Emissor                  string             `bson:"emissor"`
	EmissorLigado            string             `bson:"emissor_ligado"`
	IDDocumento              int                `bson:"id_documento"`
	Pais                     string             `bson:"pais"`
	PessoaFisicaJuridica     string             `bson:"pessoa_fisica_juridica"`
	QuantidadeAquisicoes     float64            `bson:"quantidade_aquisicoes"`
	QuantidadePosicaoFinal   float64            `bson:"quantidade_posicao_final"`
	QuantidadeVendas         float64            `bson:"quantidade_vendas"`
	RiscoEmissor             string             `bson:"risco_emissor"`
	TipoAplicacao            string             `bson:"tipo_aplicacao"`
	TipoAtivo                string             `bson:"tipo_ativo"`
	TipoFundo                string             `bson:"tipo_fundo"`
	TipoNegociacao           string             `bson:"tipo_negociacao"`
	ValorAquisicoes          float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal   float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal float64            `bson:"valor_mercado_posicao_final"`
	ValorPatrimonioLiquido   float64            `bson:"valor_patrimonio_liquido"`
	ValorVendas              float64            `bson:"valor_vendas"`
}

type CdaFiimConfidencialEntity struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	FundoCnpj                string             `bson:"cnpj_fundo"`
	DenominacaoSocial        string             `bson:"denominacao_social"`
	DataCompetencia          time.Time          `bson:"data_competencia"`
	PrazoConfidencial        time.Time          `bson:"prazo_confidencial"`
	IDDocumento              int                `bson:"id_documento"`
	TipoAplicacao            string             `bson:"tipo_aplicacao"`
	TipoFundo                string             `bson:"tipo_fundo"`
	ValorAquisicoes          float64            `bson:"valor_aquisicoes"`
	ValorCustoPosicaoFinal   float64            `bson:"valor_custo_posicao_final"`
	ValorMercadoPosicaoFinal float64            `bson:"valor_mercado_posicao_final"`
	ValorVendas              float64            `bson:"valor_vendas"`
}

type CdaPatrimonioLiquidoEntity struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	FundoCnpj         string             `bson:"cnpj_fundo"`
	DenominacaoSocial string             `bson:"denominacao_social"`
	DataCompetencia   time.Time          `bson:"data_competencia"`
	TipoFundo         string             `bson:"tipo_fundo"`
	ValorPatrimonio   float64            `bson:"valor_patrimonio"`
}
