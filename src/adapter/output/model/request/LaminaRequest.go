package request

// Fundo representa a estrutura do fundo com tags JSON e comentários descritivos
type LaminaRequest struct {
	AnoAnterExemplo           string `json:"ANO_ANTER_EXEMPLO,omitempty"`             // Ano anterior ao de referência (exemplo para comparar custos e benefícios de se investir no fundo)
	AnoExemplo                string `json:"ANO_EXEMPLO,omitempty"`                   // Ano de referência (exemplo para comparar custos e benefícios de se investir no fundo)
	AnoSemRentab              string `json:"ANO_SEM_RENTAB,omitempty"`                // Anos em que não foram apresentados dados de rentabilidade
	CalcRentabFundo           string `json:"CALC_RENTAB_FUNDO,omitempty"`             // Fórmula de cálculo de sua rentabilidade
	CalcRentabFundoGatilho    string `json:"CALC_RENTAB_FUNDO_GATILHO,omitempty"`     // Descrição da fórmula de cálculo da rentabilidade do fundo, incluindo todas as condições (gatilhos) e cláusulas que afetarão o desempenho
	ClasseRiscoAdmin          string `json:"CLASSE_RISCO_ADMIN,omitempty"`            // Risco do fundo de acordo com a classificação do administrador (escala de 1 a 5)
	CnpjFundo                 string `json:"CNPJ_FUNDO,omitempty"`                    // CNPJ do fundo
	CondicCaren               string `json:"CONDIC_CAREN,omitempty"`                  // Condições do período de carência
	CondicEntr                string `json:"CONDIC_ENTR,omitempty"`                   // Condições de entrada
	CondicSaida               string `json:"CONDIC_SAIDA,omitempty"`                  // Condições de saída
	ConflitoVenda             string `json:"CONFLITO_VENDA,omitempty"`                // Informa se existe conflito de interesses no esforço de venda
	ConversaoCotaCanc         string `json:"CONVERSAO_COTA_CANC,omitempty"`           // Evento considerado para o cálculo do valor das cotas canceladas
	ConversaoCotaCompra       string `json:"CONVERSAO_COTA_COMPRA,omitempty"`         // Evento considerado para o cálculo do valor das cotas compradas
	DenomSocial               string `json:"DENOM_SOCIAL,omitempty"`                  // Denominação Social
	DerivProtecaoCarteira     string `json:"DERIV_PROTECAO_CARTEIRA,omitempty"`       // Indica se o fundo utiliza derivativos apenas para proteção da carteira
	DistribGestorUnico        string `json:"DISTRIB_GESTOR_UNICO,omitempty"`          // Informa se o principal distribuidor oferta, para o público alvo do fundo, preponderantemente fundos geridos por um único gestor, ou por gestores ligados a um mesmo grupo econômico
	DsRentabGatilho           string `json:"DS_RENTAB_GATILHO,omitempty"`             // Descrição de como o cenário/gatilho afeta a variação do desempenho do fundo
	DtComptc                  string `json:"DT_COMPTC,omitempty"`                     // Data de competência do documento
	DtFimDespesa              string `json:"DT_FIM_DESPESA,omitempty"`                // Fim do período considerado para o cálculo das despesas
	DtIniAtiv5Ano             string `json:"DT_INI_ATIV_5ANO,omitempty"`              // Data de início do funcionamento do fundo (quando o fundo tiver sido constituído há menos de 5 anos)
	DtIniDespesa              string `json:"DT_INI_DESPESA,omitempty"`                // Início do período considerado para o cálculo das despesas
	EnderEletronico           string `json:"ENDER_ELETRONICO,omitempty"`              // Endereço eletrônico
	EnderEletronicoDespesa    string `json:"ENDER_ELETRONICO_DESPESA,omitempty"`      // Endereço eletrônico do administrador onde a descrição das despesas do fundo pode ser encontrada
	EnderEletronicoReclamacao string `json:"ENDER_ELETRONICO_RECLAMACAO,omitempty"`   // Endereço eletrônico e demais canais disponíveis para o encaminhamento de reclamações
	HoraAplicResgate          string `json:"HORA_APLIC_RESGATE,omitempty"`            // Horário para aplicação e resgate
	IndiceRefer               string `json:"INDICE_REFER,omitempty"`                  // Índice de referência do fundo
	InvestAdic                string `json:"INVEST_ADIC,omitempty"`                   // Investimento adicional
	InvestInicialMin          string `json:"INVEST_INICIAL_MIN,omitempty"`            // Investimento inicial mínimo
	NmFantasia                string `json:"NM_FANTASIA,omitempty"`                   // Nome fantasia do fundo
	Objetivo                  string `json:"OBJETIVO,omitempty"`                      // Objetivos do fundo
	PolitInvest               string `json:"POLIT_INVEST,omitempty"`                  // Política de investimento
	PrAtivoEmissor            string `json:"PR_ATIVO_EMISSOR,omitempty"`              // Limite de concentração em ativos de um só emissor que não seja a União Federal
	PrPlAlavanc               string `json:"PR_PL_ALAVANC,omitempty"`                 // Limite de alavancagem (% em relação ao PL)
	PrPlAplicMaxFundoUnico    string `json:"PR_PL_APLIC_MAX_FUNDO_UNICO,omitempty"`   // Percentual máximo do PL que pode ser aplicado em um só fundo de investimento
	PrPlAtivoCredPriv         string `json:"PR_PL_ATIVO_CRED_PRIV,omitempty"`         // Percentual do PL que pode ser aplicado em crédito privado
	PrPlAtivoExterior         string `json:"PR_PL_ATIVO_EXTERIOR,omitempty"`          // Percentual do PL que pode ser aplicado em ativos do exterior
	PrPlDespesa               string `json:"PR_PL_DESPESA,omitempty"`                 // Despesas pagas pelo fundo (em % do PL diário médio)
	PrRentabFundo5Ano         string `json:"PR_RENTAB_FUNDO_5ANO,omitempty"`          // Rentabilidade acumulada nos últimos 5 anos (valor percentual)
	PrVariacaoIndiceRefer5Ano string `json:"PR_VARIACAO_INDICE_REFER_5ANO,omitempty"` // Variação do índice de referência do fundo, nos últimos 5 anos (valor percentual)
	PrVariacaoPerfm           string `json:"PR_VARIACAO_PERFM,omitempty"`             // Variação do desempenho do fundo (valor percentual)
	PublicoAlvo               string `json:"PUBLICO_ALVO,omitempty"`                  // Público-alvo
	QtAnoPerda                string `json:"QT_ANO_PERDA,omitempty"`                  // Número de anos em que o fundo perdeu parte do patrimônio que detinha no início do respectivo ano
	QtDiaCaren                string `json:"QT_DIA_CAREN,omitempty"`                  // Número de dias de carência
	QtDiaConversaoCotaCompra  string `json:"QT_DIA_CONVERSAO_COTA_COMPRA,omitempty"`  // Número de dias para conversão de cotas contado da data da aplicação
	QtDiaConversaoCotaResgate string `json:"QT_DIA_CONVERSAO_COTA_RESGATE,omitempty"` // Número de dias para conversão de cotas após o pedido de resgate
	QtDiaPagtoResgate         string `json:"QT_DIA_PAGTO_RESGATE,omitempty"`          // Prazo para o efetivo pagamento dos resgates, contado a partir da data do pedido
	QtDiaSaida                string `json:"QT_DIA_SAIDA,omitempty"`                  // Número de dias de saída para resgatar suas cotas do fundo
	RemunDistrib              string `json:"REMUN_DISTRIB,omitempty"`                 // Descrição da forma de remuneração dos distribuidores
	RentabGatilho             string `json:"RENTAB_GATILHO,omitempty"`                // Indicação dos valores dos cenários ou gatilhos que afetam a rentabilidade
	ResgateMin                string `json:"RESGATE_MIN,omitempty"`                   // Resgate mínimo
	RestrInvest               string `json:"RESTR_INVEST,omitempty"`                  // Restrições de investimento
	RiscoPerda                string `json:"RISCO_PERDA,omitempty"`                   // Indica se as estratégias de investimento do fundo podem resultar em perdas patrimoniais significativas para seus cotistas
	RiscoPerdaNegativo        string `json:"RISCO_PERDA_NEGATIVO,omitempty"`          // Indica se as estratégias de investimento do fundo podem resultar em perdas superiores ao capital aplicado e a consequente obrigação do cotista de aportar recursos adicionais para cobrir o prejuízo do fundo
	TaxaAdm                   string `json:"TAXA_ADM,omitempty"`                      // Taxa de administração
	TaxaAdmMax                string `json:"TAXA_ADM_MAX,omitempty"`                  // Taxa de administração máxima (quando variável)
	TaxaAdmMin                string `json:"TAXA_ADM_MIN,omitempty"`                  // Taxa de administração mínima (quando variável)
	TaxaAdmObs                string `json:"TAXA_ADM_OBS,omitempty"`                  // Comentários ou esclarecimentos sobre a taxa de administração
	TaxaEntr                  string `json:"TAXA_ENTR,omitempty"`                     // Taxa de entrada
	TaxaPerfm                 string `json:"TAXA_PERFM,omitempty"`                    // Taxa de performance
	TaxaSaida                 string `json:"TAXA_SAIDA,omitempty"`                    // Taxa de saída
	TelSac                    string `json:"TEL_SAC,omitempty"`                       // Telefone do serviço de atendimento ao cotista
	TpDiaPagtoResgate         string `json:"TP_DIA_PAGTO_RESGATE,omitempty"`          // Tipo de prazo para o efetivo pagamento dos resgates
	TpTaxaAdm                 string `json:"TP_TAXA_ADM,omitempty"`                   // Tipo da taxa de administração
	VlAjustePerfmExemplo      string `json:"VL_AJUSTE_PERFM_EXEMPLO,omitempty"`       // Valor do ajuste sobre performance individual (exemplo para comparar custos e benefícios de se investir no fundo)
	VlDespesa3Ano             string `json:"VL_DESPESA_3ANO,omitempty"`               // Despesas previstas para 3 anos (se a taxa total de despesas se mantiver constante)
	VlDespesa5Ano             string `json:"VL_DESPESA_5ANO,omitempty"`               // Despesas previstas para 5 anos (se a taxa total de despesas se mantiver constante)
	VlDespesaExemplo          string `json:"VL_DESPESA_EXEMPLO,omitempty"`            // Valor total das despesas do fundo (exemplo para comparar custos e benefícios de se investir no fundo)
	VlImpostoExemplo          string `json:"VL_IMPOSTO_EXEMPLO,omitempty"`            // Valor dos impostos que seriam pagos (exemplo para comparar custos e benefícios de se investir no fundo)
	VlMinPerman               string `json:"VL_MIN_PERMAN,omitempty"`                 // Valor mínimo para permanência
	VlPatrimLiq               string `json:"VL_PATRIM_LIQ,omitempty"`                 // Valor do patrimônio líquido
	VlResgateExemplo          string `json:"VL_RESGATE_EXEMPLO,omitempty"`            // Valor que poderia ser resgatado já deduzido dos impostos (exemplo para comparar custos e benefícios de se investir no fundo)
	VlRetorno3Ano             string `json:"VL_RETORNO_3ANO,omitempty"`               // Retorno bruto hipotético para 3 anos após dedução das despesas e do valor do investimento original (antes da incidência de impostos, de taxas de ingresso e/ou saída, ou de taxa de performance)
	VlRetorno5Ano             string `json:"VL_RETORNO_5ANO,omitempty"`               // Retorno bruto hipotético para 5 anos após dedução das despesas e do valor do investimento original (antes da incidência de impostos, de taxas de ingresso e/ou saída, ou de taxa de performance)
	VlTaxaEntrExemplo         string `json:"VL_TAXA_ENTR_EXEMPLO,omitempty"`          // Valor da taxa de ingresso (exemplo para comparar custos e benefícios de se investir no fundo)
	VlTaxaSaidaExemplo        string `json:"VL_TAXA_SAIDA_EXEMPLO,omitempty"`         // Valor da taxa de saída (exemplo para comparar custos e benefícios de se investir no fundo)
	InfSac                    string `json:"INF_SAC,omitempty"`                       //
}

type LaminaCarteiraRequest struct {
	CnpjFundo   string `json:"CNPJ_FUNDO,omitempty"`   // CNPJ do fundo
	DenomSocial string `json:"DENOM_SOCIAL,omitempty"` // Denominação Social
	DtComptc    string `json:"DT_COMPTC,omitempty"`    // Data de competência do documento (formato: "AAAA-MM-DD")
	PrPlAtivo   string `json:"PR_PL_ATIVO,omitempty"`  // Concentração dos investimentos do fundo por tipo de ativo (em % do PL)
	TpAtivo     string `json:"TP_ATIVO,omitempty"`     // Tipo de ativo
}

type LaminaRentabilidadeAnoRequest struct {
	AnoReferencia            string `json:"ANO_RENTAB,omitempty"`                   // Ano referente à rentabilidade do fundo
	CnpjFundo                string `json:"CNPJ_FUNDO,omitempty"`                   // CNPJ do fundo
	DenomSocial              string `json:"DENOM_SOCIAL,omitempty"`                 // Denominação Social
	DtComptc                 string `json:"DT_COMPTC,omitempty"`                    // Data de competência do documento (formato: "AAAA-MM-DD")
	PrPerfIndiceReferAno     string `json:"PR_PERFM_INDICE_REFER_ANO,omitempty"`    // Desempenho anual do fundo como percentual do índice de referência
	PrRentabAno              string `json:"PR_RENTAB_ANO,omitempty"`                // Rentabilidade do fundo no ano de referência (valor percentual)
	PrVariacaoIndiceReferAno string `json:"PR_VARIACAO_INDICE_REFER_ANO,omitempty"` // Variação anual do índice de referência do fundo (valor percentual)
	RentabAnoObservacoes     string `json:"RENTAB_ANO_OBS,omitempty"`               // Comentários sobre a rentabilidade do fundo no ano de referência
}

type LaminaRentabilidadeMesRequest struct {
	CnpjFundo                string `json:"CNPJ_FUNDO,omitempty"`                   // CNPJ do fundo
	DenomSocial              string `json:"DENOM_SOCIAL,omitempty"`                 // Denominação Social
	DtComptc                 string `json:"DT_COMPTC,omitempty"`                    // Data de competência do documento (formato: "AAAA-MM-DD")
	MesReferencia            string `json:"MES_RENTAB,omitempty"`                   // Mês referente à rentabilidade do fundo
	PrPerfIndiceReferMes     string `json:"PR_PERFM_INDICE_REFER_MES,omitempty"`    // Desempenho mensal do fundo como percentual do índice de referência
	PrRentabMes              string `json:"PR_RENTAB_MES,omitempty"`                // Rentabilidade do fundo no mês de referência (valor percentual)
	PrVariacaoIndiceReferMes string `json:"PR_VARIACAO_INDICE_REFER_MES,omitempty"` // Variação mensal do índice de referência do fundo (valor percentual)
	RentabMesObservacoes     string `json:"RENTAB_MES_OBS,omitempty"`               // Comentários sobre a rentabilidade do fundo no mês de referência
}
