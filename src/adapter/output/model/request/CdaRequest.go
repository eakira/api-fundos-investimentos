package request

type CdaSelicRequest struct {
	CodigoConta              string  `json:"CD_CONTA_BALCTE,omitempty"`    // Código da conta balancete
	FundoCnpj                Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo
	DataCompetencia          Date    `json:"DT_COMPTC,omitempty"`          // Data de competência
	PlanoContabil            string  `json:"PLANO_CONTA_BALCTE,omitempty"` // Plano contábil
	SaldoConta               string  `json:"VL_SALDO_BALCTE,omitempty"`    // Saldo da conta balancete
	CdIsin                   string  `json:"CD_ISIN,omitempty"`            // Código ISIN (International Securities Identification Number)
	CdSelic                  string  `json:"CD_SELIC,omitempty"`           // Código SELIC
	DenominacaoSocial        string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social
	DataEmissao              Date    `json:"DT_EMISSAO,omitempty"`         // Data de emissão (AAAA-MM-DD)
	DataVencimento           Date    `json:"DT_VENC,omitempty"`            // Data de vencimento (AAAA-MM-DD)
	DataCompetenciaDoc       Date    `json:"DT_COMPTC_DOC,omitempty"`      // Data de competência do documento
	PrazoConfidencialidade   Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	EmissorLigado            string  `json:"EMISSOR_LIGADO,omitempty"`     // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	QuantidadeAquisicoes     Decimal `json:"QT_AQUIS_NEGOC,omitempty"`     // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal   Decimal `json:"QT_POS_FINAL,omitempty"`       // Quantidade da posição final
	QuantidadeVendas         Decimal `json:"QT_VENDA_NEGOC,omitempty"`     // Quantidade de vendas dos negócios realizados no mês
	TipoAplicacao            string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação
	TipoAtivo                string  `json:"TP_ATIVO,omitempty"`           // Tipo de ativo
	TipoFundo                string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo
	TipoNegociacao           string  `json:"TP_NEGOC,omitempty"`           // Tipo de negociação (Para negociação/Mantido até o vencimento)
	TipoTituloPublico        string  `json:"TP_TITPUB,omitempty"`          // Tipo de título público
	ValorAquisicoes          Decimal `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal   Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal Decimal `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas              Decimal `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaFundosInvestimentosRequest struct {
	FundoCnpj                Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo (20 caracteres)
	FundoCnpjCota            Cnpj    `json:"CNPJ_FUNDO_COTA,omitempty"`    // CNPJ do fundo investido (20 caracteres)
	DenominacaoSocial        string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social (100 caracteres)
	DataCompetencia          Date    `json:"DT_COMPTC,omitempty"`          // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencialidade   Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	EmissorLigado            string  `json:"EMISSOR_LIGADO,omitempty"`     // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	NomeFundoCota            string  `json:"NM_FUNDO_COTA,omitempty"`      // Denominação social do fundo investido (100 caracteres)
	QuantidadeAquisicoes     Decimal `json:"QT_AQUIS_NEGOC,omitempty"`     // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal   Decimal `json:"QT_POS_FINAL,omitempty"`       // Quantidade da posição final
	QuantidadeVendas         Decimal `json:"QT_VENDA_NEGOC,omitempty"`     // Quantidade de vendas dos negócios realizados no mês
	TipoAplicacao            string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação (150 caracteres)
	TipoAtivo                string  `json:"TP_ATIVO,omitempty"`           // Tipo de ativo (150 caracteres)
	TipoFundo                string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo (15 caracteres)
	TipoNegociacao           string  `json:"TP_NEGOC,omitempty"`           // Tipo de negociação (Para negociação/Mantido até o vencimento) (24 caracteres)
	ValorAquisicoes          Decimal `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal   Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal Decimal `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas              Decimal `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaSwapRequest struct {
	CodigoSwap               string  `json:"CD_SWAP,omitempty"`            // Código SWAP (50 caracteres)
	FundoCnpj                Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo (20 caracteres)
	DenominacaoSocial        string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social (100 caracteres)
	DescricaoTipoAtivo       string  `json:"DS_SWAP,omitempty"`            // Descrição do tipo de ativo SWAP (100 caracteres)
	DataCompetencia          Date    `json:"DT_COMPTC,omitempty"`          // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencialidade   Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	EmissorLigado            string  `json:"EMISSOR_LIGADO,omitempty"`     // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	QuantidadeAquisicoes     Decimal `json:"QT_AQUIS_NEGOC,omitempty"`     // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal   Decimal `json:"QT_POS_FINAL,omitempty"`       // Quantidade da posição final
	QuantidadeVendas         Decimal `json:"QT_VENDA_NEGOC,omitempty"`     // Quantidade de vendas dos negócios realizados no mês
	TipoAplicacao            string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação (150 caracteres)
	TipoAtivo                string  `json:"TP_ATIVO,omitempty"`           // Tipo de ativo (150 caracteres)
	TipoFundo                string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo (15 caracteres)
	TipoNegociacao           string  `json:"TP_NEGOC,omitempty"`           // Tipo de negociação (Para negociação/Mantido até o vencimento) (24 caracteres)
	ValorAquisicoes          Decimal `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal   Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal Decimal `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas              Decimal `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaDemaisAtivosRequest struct {
	CodigoAtivo              string  `json:"CD_ATIVO,omitempty"`           // Código do ativo (100 caracteres)
	CodigoISIN               string  `json:"CD_ISIN,omitempty"`            // Código ISIN (International Securities Identification Number) (12 caracteres)
	FundoCnpj                Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo (20 caracteres)
	DenominacaoSocial        string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social (100 caracteres)
	DescricaoAtivo           string  `json:"DS_ATIVO,omitempty"`           // Descrição do ativo (100 caracteres)
	DataCompetencia          Date    `json:"DT_COMPTC,omitempty"`          // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencialidade   Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	DataFimVigencia          Date    `json:"DT_FIM_VIGENCIA,omitempty"`    // Data fim da vigência (AAAA-MM-DD)
	DataInicioVigencia       Date    `json:"DT_INI_VIGENCIA,omitempty"`    // Data início da vigência (AAAA-MM-DD)
	EmissorLigado            string  `json:"EMISSOR_LIGADO,omitempty"`     // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	QuantidadeAquisicoes     Decimal `json:"QT_AQUIS_NEGOC,omitempty"`     // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal   Decimal `json:"QT_POS_FINAL,omitempty"`       // Quantidade da posição final
	QuantidadeVendas         Decimal `json:"QT_VENDA_NEGOC,omitempty"`     // Quantidade de vendas dos negócios realizados no mês
	TipoAplicacao            string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação (150 caracteres)
	TipoAtivo                string  `json:"TP_ATIVO,omitempty"`           // Tipo de ativo (150 caracteres)
	TipoFundo                string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo (15 caracteres)
	TipoNegociacao           string  `json:"TP_NEGOC,omitempty"`           // Tipo de negociação (Para negociação/Mantido até o vencimento) (24 caracteres)
	ValorAquisicoes          Decimal `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal   Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal Decimal `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas              Decimal `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaDepositoAPrazoOutrosAtivosRequest struct {
	NomeAgenciaClassificacaoRisco string  `json:"AG_RISCO,omitempty"`           // Nome da agência de classificação de risco (200 caracteres)
	CodigoIndexadorPosFixados     string  `json:"CD_INDEXADOR_POSFX,omitempty"` // Código do indexador (somente pós-fixados) (50 caracteres)
	CnpjEmissor                   Cnpj    `json:"CNPJ_EMISSOR,omitempty"`       // CNPJ do emissor (20 caracteres)
	FundoCnpj                     Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo (20 caracteres)
	DenominacaoSocial             string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social (100 caracteres)
	DescricaoIndexadorPosFixados  string  `json:"DS_INDEXADOR_POSFX,omitempty"` // Descrição do indexador (somente pós-fixados) (100 caracteres)
	DataCompetencia               Date    `json:"DT_COMPTC,omitempty"`          // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencialidade        Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	DataClassificacaoRisco        Date    `json:"DT_RISCO,omitempty"`           // Data da classificação de risco (AAAA-MM-DD)
	DataVencimento                Date    `json:"DT_VENC,omitempty"`            // Data de vencimento (AAAA-MM-DD)
	NomeEmissor                   string  `json:"EMISSOR,omitempty"`            // Nome do Emissor (200 caracteres)
	EmissorLigado                 string  `json:"EMISSOR_LIGADO,omitempty"`     // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	GrauRisco                     string  `json:"GRAU_RISCO,omitempty"`         // Grau de risco atribuído (6 caracteres)
	PercentualCupomPosFixados     Decimal `json:"PR_CUPOM_POSFX,omitempty"`     // Percentual do cupom (somente pós-fixados)
	PercentualIndexadorPosFixados Decimal `json:"PR_INDEXADOR_POSFX,omitempty"` // Percentual do indexador (somente pós-fixados)
	PercentualTaxaPrefixada       Decimal `json:"PR_TAXA_PREFX,omitempty"`      // Percentual da taxa contratada (somente pré-fixados)
	QuantidadeAquisicoes          Decimal `json:"QT_AQUIS_NEGOC,omitempty"`     // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal        Decimal `json:"QT_POS_FINAL,omitempty"`       // Quantidade da posição final
	QuantidadeVendas              Decimal `json:"QT_VENDA_NEGOC,omitempty"`     // Quantidade de vendas dos negócios realizados no mês
	RiscoEmissor                  string  `json:"RISCO_EMISSOR,omitempty"`      // Indica se o emissor do título possui classificação de risco (S/N)
	TituloPosFixado               string  `json:"TITULO_POSFX,omitempty"`       // Indica se é título pós-fixado (S/N)
	TipoAplicacao                 string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação (150 caracteres)
	TipoAtivo                     string  `json:"TP_ATIVO,omitempty"`           // Tipo de ativo (150 caracteres)
	TipoFundo                     string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo (15 caracteres)
	TipoNegociacao                string  `json:"TP_NEGOC,omitempty"`           // Tipo de negociação (Para negociação/Mantido até o vencimento) (24 caracteres)
	ValorAquisicoes               Decimal `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal        Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal      Decimal `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas                   Decimal `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaAgroCreditoPrivadoRequest struct {
	CodigoIndexadorPosFixados      string  `json:"CD_INDEXADOR_POSFX,omitempty"`            // Código do indexador (somente pós-fixados) (50 caracteres)
	FundoCnpj                      Cnpj    `json:"CNPJ_FUNDO,omitempty"`                    // CNPJ do fundo (20 caracteres)
	CnpjInstituicaoFinanceiraCoobr Cnpj    `json:"CNPJ_INSTITUICAO_FINANC_COOBR,omitempty"` // CNPJ da instituição financeira com coobrigação (20 caracteres)
	CpfCnpjEmissor                 Cnpj    `json:"CPF_CNPJ_EMISSOR,omitempty"`              // Código de identificação do emissor, pessoa física ou jurídica (20 caracteres)
	DenominacaoSocial              string  `json:"DENOM_SOCIAL,omitempty"`                  // Denominação Social (100 caracteres)
	DescricaoIndexadorPosFixados   string  `json:"DS_INDEXADOR_POSFX,omitempty"`            // Descrição do indexador (somente pós-fixados) (100 caracteres)
	DataCompetencia                Date    `json:"DT_COMPTC,omitempty"`                     // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencialidade         Date    `json:"DT_CONFID_APLIC,omitempty"`               // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	DataVencimento                 Date    `json:"DT_VENC,omitempty"`                       // Data de vencimento (AAAA-MM-DD)
	NomeEmissor                    string  `json:"EMISSOR,omitempty"`                       // Nome do Emissor (200 caracteres)
	EmissorLigado                  string  `json:"EMISSOR_LIGADO,omitempty"`                // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	PessoaFisicaJuridicaEmissor    string  `json:"PF_PJ_EMISSOR,omitempty"`                 // Indica se o emissor é pessoa física ou jurídica (PF/PJ)
	PercentualCupomPosFixados      Decimal `json:"PR_CUPOM_POSFX,omitempty"`                // Percentual do cupom (somente pós-fixados)
	PercentualIndexadorPosFixados  Decimal `json:"PR_INDEXADOR_POSFX,omitempty"`            // Percentual do indexador (somente pós-fixados)
	PercentualTaxaPrefixada        Decimal `json:"PR_TAXA_PREFX,omitempty"`                 // Percentual da taxa contratada (somente pré-fixados)
	QuantidadeAquisicoes           Decimal `json:"QT_AQUIS_NEGOC,omitempty"`                // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal         Decimal `json:"QT_POS_FINAL,omitempty"`                  // Quantidade da posição final
	QuantidadeVendas               Decimal `json:"QT_VENDA_NEGOC,omitempty"`                // Quantidade de vendas dos negócios realizados no mês
	TituloRegistradoCetip          string  `json:"TITULO_CETIP,omitempty"`                  // Indica se é título registrado na CETIP (S/N)
	TituloPossuiGarantia           string  `json:"TITULO_GARANTIA,omitempty"`               // Indica se o título possui garantia ou seguro (S/N)
	TituloPosFixado                string  `json:"TITULO_POSFX,omitempty"`                  // Indica se é título pós-fixado (S/N)
	TipoAplicacao                  string  `json:"TP_APLIC,omitempty"`                      // Tipo de aplicação (150 caracteres)
	TipoAtivo                      string  `json:"TP_ATIVO,omitempty"`                      // Tipo de ativo (150 caracteres)
	TipoFundo                      string  `json:"TP_FUNDO,omitempty"`                      // Tipo de fundo (15 caracteres)
	TipoNegociacao                 string  `json:"TP_NEGOC,omitempty"`                      // Tipo de negociação (Para negociação/Mantido até o vencimento) (24 caracteres)
	ValorAquisicoes                Decimal `json:"VL_AQUIS_NEGOC,omitempty"`                // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal         Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"`            // Valor de custo da posição final
	ValorMercadoPosicaoFinal       Decimal `json:"VL_MERC_POS_FINAL,omitempty"`             // Valor de mercado da posição final
	ValorVendas                    Decimal `json:"VL_VENDA_NEGOC,omitempty"`                // Valor das vendas dos negócios realizados no mês
}

type CdaInvestimentosExteriorRequest struct {
	AgenciaRisco                      string  `json:"AG_RISCO,omitempty"`               // Nome da agência de classificação de risco (200 caracteres)
	BolsaMercado                      string  `json:"BV_MERC,omitempty"`                // Bolsa ou Mercado de balcão (100 caracteres)
	CodigoAtivoBolsaMercado           string  `json:"CD_ATIVO_BV_MERC,omitempty"`       // Código do ativo na Bolsa ou Mercado de balcão onde foi adquirido (12 caracteres)
	CodigoBolsaMercado                string  `json:"CD_BV_MERC,omitempty"`             // Código da Bolsa ou Mercado de balcão (6 caracteres)
	CodigoPais                        string  `json:"CD_PAIS,omitempty"`                // Código do país (3 caracteres)
	FundoCnpj                         Cnpj    `json:"CNPJ_FUNDO,omitempty"`             // CNPJ do fundo (20 caracteres)
	DenominacaoSocial                 string  `json:"DENOM_SOCIAL,omitempty"`           // Denominação Social (100 caracteres)
	DescricaoAtivoExterior            string  `json:"DS_ATIVO_EXTERIOR,omitempty"`      // Descrição do ativo no exterior (50 caracteres)
	DataCompetencia                   Date    `json:"DT_COMPTC,omitempty"`              // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencialidade            Date    `json:"DT_CONFID_APLIC,omitempty"`        // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	DataClassificacaoRisco            Date    `json:"DT_RISCO,omitempty"`               // Data da classificação de risco (AAAA-MM-DD)
	DataVencimento                    Date    `json:"DT_VENC,omitempty"`                // Data de vencimento (AAAA-MM-DD)
	NomeEmissor                       string  `json:"EMISSOR,omitempty"`                // Nome do Emissor (200 caracteres)
	EmissorLigado                     string  `json:"EMISSOR_LIGADO,omitempty"`         // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	GrauRisco                         string  `json:"GRAU_RISCO,omitempty"`             // Grau de risco atribuído (6 caracteres)
	InvestimentoColetivo              string  `json:"INVEST_COLETIVO,omitempty"`        // Indica se é veículo de investimento coletivo (S/N)
	GestaoVeiculoInvestimentoColetivo string  `json:"INVEST_COLETIVO_GESTOR,omitempty"` // Indica se a gestão da carteira do veículo de investimento coletivo conta com influência, direta ou indireta, do gestor (S/N)
	Pais                              string  `json:"PAIS,omitempty"`                   // País (100 caracteres)
	QuantidadeAquisicoes              Decimal `json:"QT_AQUIS_NEGOC,omitempty"`         // Quantidade de aquisições dos negócios realizados no mês
	QuantidadeAtivosExterior          Decimal `json:"QT_ATIVO_EXTERIOR,omitempty"`      // Quantidade de ativos no exterior
	QuantidadePosicaoFinal            Decimal `json:"QT_POS_FINAL,omitempty"`           // Quantidade da posição final
	QuantidadeVendas                  Decimal `json:"QT_VENDA_NEGOC,omitempty"`         // Quantidade de vendas dos negócios realizados no mês
	RiscoEmissor                      string  `json:"RISCO_EMISSOR,omitempty"`          // Indica se o emissor do título possui classificação de risco (S/N)
	TipoAplicacao                     string  `json:"TP_APLIC,omitempty"`               // Tipo de aplicação (150 caracteres)
	TipoAtivo                         string  `json:"TP_ATIVO,omitempty"`               // Tipo de ativo (150 caracteres)
	TipoFundo                         string  `json:"TP_FUNDO,omitempty"`               // Tipo de fundo (15 caracteres)
	TipoNegociacao                    string  `json:"TP_NEGOC,omitempty"`               // Tipo de negociação (Para negociação/Mantido até o vencimento) (24 caracteres)
	ValorAquisicoes                   Decimal `json:"VL_AQUIS_NEGOC,omitempty"`         // Valor das aquisições dos negócios realizados no mês
	ValorAtivoExterior                Decimal `json:"VL_ATIVO_EXTERIOR,omitempty"`      // Valor do ativo no exterior
	ValorCustoPosicaoFinal            Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"`     // Valor de custo da posição final
	ValorMercadoPosicaoFinal          Decimal `json:"VL_MERC_POS_FINAL,omitempty"`      // Valor de mercado da posição final
	ValorVendas                       Decimal `json:"VL_VENDA_NEGOC,omitempty"`         // Valor das vendas dos negócios realizados no mês
}

// DetalhesInvestimento representa informações detalhadas sobre um investimento.
type CdaDemaisAtivosNaoCodificadosRequest struct {
	FundoCnpj                   Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo (20 caracteres)
	CpfCnpjEmissor              Cnpj    `json:"CPF_CNPJ_EMISSOR,omitempty"`   // Código de identificação do emissor, pessoa física ou jurídica (20 caracteres)
	DenominacaoSocial           string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social (100 caracteres)
	DescricaoAtivo              string  `json:"DS_ATIVO,omitempty"`           // Descrição do ativo (1000 caracteres)
	DataCompetencia             Date    `json:"DT_COMPTC,omitempty"`          // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencialidade      Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	NomeEmissor                 string  `json:"EMISSOR,omitempty"`            // Nome do Emissor (150 caracteres)
	EmissorLigado               string  `json:"EMISSOR_LIGADO,omitempty"`     // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (S/N)
	PessoaFisicaJuridicaEmissor string  `json:"PF_PJ_EMISSOR,omitempty"`      // Indica se o emissor é pessoa física ou jurídica (PF/PJ)
	QuantidadeAquisicoes        Decimal `json:"QT_AQUIS_NEGOC,omitempty"`     // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal      Decimal `json:"QT_POS_FINAL,omitempty"`       // Quantidade da posição final
	QuantidadeVendas            Decimal `json:"QT_VENDA_NEGOC,omitempty"`     // Quantidade de vendas dos negócios realizados no mês
	TipoAplicacao               string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação (150 caracteres)
	TipoAtivo                   string  `json:"TP_ATIVO,omitempty"`           // Tipo de ativo (150 caracteres)
	TipoFundo                   string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo (15 caracteres)
	TipoNegociacao              string  `json:"TP_NEGOC,omitempty"`           // Tipo de negociação (Para negociação/Mantido até o vencimento) (24 caracteres)
	ValorAquisicoes             Decimal `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal      Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal    Decimal `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas                 Decimal `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaConfidencialRequest struct {
	FundoCnpj           Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo (20 caracteres)
	DenominacaoSocial   string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social (100 caracteres)
	DataCompetencia     Date    `json:"DT_COMPTC,omitempty"`          // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencial   Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	TipoAplicacao       string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação (150 caracteres)
	TipoFundo           string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo (15 caracteres)
	ValorAquisicoes     float64 `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicao   float64 `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicao float64 `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas         float64 `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaFiimRequest struct {
	BolsaMercado             string  `json:"BV_MERC,omitempty"`            // Bolsa ou Mercado de balcão (47 caracteres)
	CodigoAtivo              string  `json:"CD_ATIVO,omitempty"`           // Código do ativo (255 caracteres)
	CodigoBolsaMercado       string  `json:"CD_BV_MERC,omitempty"`         // Código da Bolsa ou Mercado de balcão (5 caracteres)
	CodigoPais               Integer `json:"CD_PAIS,omitempty"`            // Código do país
	CodigoSELIC              string  `json:"CD_SELIC,omitempty"`           // Código SELIC (50 caracteres)
	FundoCnpj                Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo (20 caracteres)
	CpfCnpjEmissor           Cnpj    `json:"CPF_CNPJ_EMISSOR,omitempty"`   // Código de identificação do emissor, pessoa física ou jurídica
	DenominacaoSocial        string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social (255 caracteres)
	DescricaoAtivo           string  `json:"DS_ATIVO,omitempty"`           // Descrição do ativo (255 caracteres)
	DataCompetencia          Date    `json:"DT_COMPTC,omitempty"`          // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencial        Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	DataInicioVigencia       Date    `json:"DT_INI_VIGENCIA,omitempty"`    // Data início da vigência (AAAA-MM-DD)
	DataVencimento           Date    `json:"DT_VENC,omitempty"`            // Data de vencimento (AAAA-MM-DD)
	Emissor                  string  `json:"EMISSOR,omitempty"`            // Nome do Emissor (255 caracteres)
	EmissorLigado            string  `json:"EMISSOR_LIGADO,omitempty"`     // Indica se o emissor da aplicação é ligado ao gestor ou administrador do fundo de investimento (1 caractere)
	IDDocumento              Integer `json:"ID_DOC,omitempty"`             // Identificador do documento
	Pais                     string  `json:"PAIS,omitempty"`               // País (25 caracteres)
	PessoaFisicaJuridica     string  `json:"PF_PJ_EMISSOR,omitempty"`      // Indica se o emissor é pessoa física ou jurídica (2 caracteres)
	QuantidadeAquisicoes     Decimal `json:"QT_AQUIS_NEGOC,omitempty"`     // Quantidade de aquisições dos negócios realizados no mês
	QuantidadePosicaoFinal   Decimal `json:"QT_POS_FINAL,omitempty"`       // Quantidade da posição final
	QuantidadeVendas         Decimal `json:"QT_VENDA_NEGOC,omitempty"`     // Quantidade de vendas dos negócios realizados no mês
	RiscoEmissor             string  `json:"RISCO_EMISSOR,omitempty"`      // Indica se o emissor do título possui classificação de risco (1 caractere)
	TipoAplicacao            string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação (8000 caracteres)
	TipoAtivo                string  `json:"TP_ATIVO,omitempty"`           // Tipo de ativo (53 caracteres)
	TipoFundo                string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo (4 caracteres)
	TipoNegociacao           string  `json:"TP_NEGOC,omitempty"`           // Tipo de negociação (24 caracteres)
	ValorAquisicoes          Decimal `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal   Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal Decimal `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorPatrimonioLiquido   Decimal `json:"VL_PATRIM_LIQ,omitempty"`      // Valor do patrimônio líquido
	ValorVendas              Decimal `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaFiimConfidencialRequest struct {
	FundoCnpj                Cnpj    `json:"CNPJ_FUNDO,omitempty"`         // CNPJ do fundo (20 caracteres)
	DenominacaoSocial        string  `json:"DENOM_SOCIAL,omitempty"`       // Denominação Social (255 caracteres)
	DataCompetencia          Date    `json:"DT_COMPTC,omitempty"`          // Data de competência do documento (AAAA-MM-DD)
	PrazoConfidencial        Date    `json:"DT_CONFID_APLIC,omitempty"`    // Prazo de confidencialidade da aplicação (AAAA-MM-DD)
	IDDocumento              Integer `json:"ID_DOC,omitempty"`             // Identificador do documento
	TipoAplicacao            string  `json:"TP_APLIC,omitempty"`           // Tipo de aplicação (8000 caracteres)
	TipoFundo                string  `json:"TP_FUNDO,omitempty"`           // Tipo de fundo (4 caracteres)
	ValorAquisicoes          Decimal `json:"VL_AQUIS_NEGOC,omitempty"`     // Valor das aquisições dos negócios realizados no mês
	ValorCustoPosicaoFinal   Decimal `json:"VL_CUSTO_POS_FINAL,omitempty"` // Valor de custo da posição final
	ValorMercadoPosicaoFinal Decimal `json:"VL_MERC_POS_FINAL,omitempty"`  // Valor de mercado da posição final
	ValorVendas              Decimal `json:"VL_VENDA_NEGOC,omitempty"`     // Valor das vendas dos negócios realizados no mês
}

type CdaPatrimonioLiquidoRequest struct {
	FundoCnpj         Cnpj    `json:"CNPJ_FUNDO,omitempty"`    // CNPJ do fundo (20 caracteres)
	DenominacaoSocial string  `json:"DENOM_SOCIAL,omitempty"`  // Denominação Social (100 caracteres)
	DataCompetencia   Date    `json:"DT_COMPTC,omitempty"`     // Data de competência do documento (AAAA-MM-DD)
	TipoFundo         string  `json:"TP_FUNDO,omitempty"`      // Tipo de fundo (15 caracteres)
	ValorPatrimonio   Decimal `json:"VL_PATRIM_LIQ,omitempty"` // Valor do patrimônio líquido (com precisão de 23 dígitos, 2 decimais)
}
