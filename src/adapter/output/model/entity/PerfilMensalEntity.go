package entity

type PerfilMensalEntity struct {
	CenarioFPRCupom                    string `bson:"cenario_fpr_cupom"`                      // Cenário de fator de risco para cupom
	CenarioFPRDolar                    string `bson:"cenario_fpr_dolar"`                      // Cenário de fator de risco para dólar
	CenarioFPRIBovespa                 string `bson:"cenario_fpr_ibovespa"`                   // Cenário de fator de risco para Ibovespa
	CenarioFPRJuros                    string `bson:"cenario_fpr_juros"`                      // Cenário de fator de risco para juros
	CenarioFPROutro                    string `bson:"cenario_fpr_outro"`                      // Outro cenário de fator de risco
	FundoCnpj                          string `bson:"cnpj_fundo"`                             // CNPJ do fundo
	ComitenteLigado1                   string `bson:"comitente_ligado_1"`                     // Comitente ligado 1
	ComitenteLigado2                   string `bson:"comitente_ligado_2"`                     // Comitente ligado 2
	ComitenteLigado3                   string `bson:"comitente_ligado_3"`                     // Comitente ligado 3
	CPFCNPJComitente1                  string `bson:"cpf_cnpj_comitente_1"`                   // CPF/CNPJ do comitente 1
	CPFCNPJComitente2                  string `bson:"cpf_cnpj_comitente_2"`                   // CPF/CNPJ do comitente 2
	CPFCNPJComitente3                  string `bson:"cpf_cnpj_comitente_3"`                   // CPF/CNPJ do comitente 3
	CPFCNPJEmissor1                    string `bson:"cpf_cnpj_emissor_1"`                     // CPF/CNPJ do emissor 1
	CPFCNPJEmissor2                    string `bson:"cpf_cnpj_emissor_2"`                     // CPF/CNPJ do emissor 2
	CPFCNPJEmissor3                    string `bson:"cpf_cnpj_emissor_3"`                     // CPF/CNPJ do emissor 3
	DelibAssemb                        string `bson:"delib_assemb"`                           // Deliberação da assembleia
	DenomSocial                        string `bson:"denom_social"`                           // Denominação social
	DataComptc                         string `bson:"dt_comptc"`                              // Data de competência
	DataCotaTaxaPerfm                  string `bson:"dt_cota_taxa_perfm"`                     // Data da cota de taxa de performance
	EmissorLigado1                     string `bson:"emissor_ligado_1"`                       // Emissor ligado 1
	EmissorLigado2                     string `bson:"emissor_ligado_2"`                       // Emissor ligado 2
	EmissorLigado3                     string `bson:"emissor_ligado_3"`                       // Emissor ligado 3
	FatorRiscoNocional                 string `bson:"fator_risco_nocional"`                   // Fator de risco nocional
	FatorRiscoOutro                    string `bson:"fator_risco_outro"`                      // Outro fator de risco
	FPR                                string `bson:"fpr"`                                    // FPR
	JustifVotoAdminAssemb              string `bson:"justif_voto_admin_assemb"`               // Justificativa de voto na assembleia administrativa
	ModVar                             string `bson:"mod_var"`                                // Modelo variável
	NrCotstBanco                       string `bson:"nr_cotst_banco"`                         // Número de cotistas de banco
	NrCotstCapitaliz                   string `bson:"nr_cotst_capitaliz"`                     // Número de cotistas de capitalização
	NrCotstCorretoraDistrib            string `bson:"nr_cotst_corretora_distrib"`             // Número de cotistas de corretora distribuidora
	NrCotstDistrib                     string `bson:"nr_cotst_distrib"`                       // Número de cotistas distribuidores
	NrCotstEAPC                        string `bson:"nr_cotst_eapc"`                          // Número de cotistas EAPC
	NrCotstEFPC                        string `bson:"nr_cotst_efpc"`                          // Número de cotistas EFPC
	NrCotstEntidPrevidCompl            string `bson:"nr_cotst_entid_previd_compl"`            // Número de cotistas entidade de previdência complementar
	NrCotstFIClube                     string `bson:"nr_cotst_fi_clube"`                      // Número de cotistas de FIC clube
	NrCotstInvnr                       string `bson:"nr_cotst_invnr"`                         // Número de cotistas investidores não residentes
	NrCotstOutro                       string `bson:"nr_cotst_outro"`                         // Número de cotistas de outro tipo
	NrCotstPFPB                        string `bson:"nr_cotst_pf_pb"`                         // Número de cotistas de pessoa física PB
	NrCotstPFVarejo                    string `bson:"nr_cotst_pf_varejo"`                     // Número de cotistas de pessoa física varejo
	NrCotstPJFinanc                    string `bson:"nr_cotst_pj_financ"`                     // Número de cotistas de pessoa jurídica financeira
	NrCotstPJNaoFinancPB               string `bson:"nr_cotst_pj_nao_financ_pb"`              // Número de cotistas de pessoa jurídica não financeira PB
	NrCotstPJNaoFinancVarejo           string `bson:"nr_cotst_pj_nao_financ_varejo"`          // Número de cotistas de pessoa jurídica não financeira varejo
	NrCotstRPPS                        string `bson:"nr_cotst_rpps"`                          // Número de cotistas RPPS
	NrCotstSegur                       string `bson:"nr_cotst_segur"`                         // Número de cotistas de seguradora
	NrDiaCemPerc                       string `bson:"nr_dia_cem_perc"`                        // Número de dias 100% de aplicação
	NrDiaCinquPerc                     string `bson:"nr_dia_cinqu_perc"`                      // Número de dias 50% de aplicação
	PFPJComitente1                     string `bson:"pf_pj_comitente_1"`                      // PF/PJ comitente 1
	PFPJComitente2                     string `bson:"pf_pj_comitente_2"`                      // PF/PJ comitente 2
	PFPJComitente3                     string `bson:"pf_pj_comitente_3"`                      // PF/PJ comitente 3
	QtdAcoesEmitidasMaiorQtdAutoriz    string `bson:"qtd_acoes_emitidas_maior_qtd_autoriz"`   // Quantidade de ações emitidas maior que quantidade autorizada
	QtdCotstEmit                       string `bson:"qtd_cotst_emit"`                         // Quantidade de cotistas emitidas
	QtdCotstResgatadas                 string `bson:"qtd_cotst_resgatadas"`                   // Quantidade de cotistas resgatadas
	QtdDiasCorridos                    string `bson:"qtd_dias_corridos"`                      // Quantidade de dias corridos
	QtdEmissaoResgate                  string `bson:"qtd_emissao_resgate"`                    // Quantidade de emissão/resgate
	QtdEmissaoResgateAgrupado          string `bson:"qtd_emissao_resgate_agrupado"`           // Quantidade de emissão/resgate agrupado
	QtdRecolhimentoIntegralizacao      string `bson:"qtd_recolhimento_integralizacao"`        // Quantidade de recolhimento para integralização
	QtdReembolsoDividendosJSCP         string `bson:"qtd_reembolso_dividendos_jscp"`          // Quantidade de reembolso de dividendos JSCP
	QtdReembolsoDividendosJSCPAgrupado string `bson:"qtd_reembolso_dividendos_jscp_agrupado"` // Quantidade de reembolso de dividendos JSCP agrupado
	QtdReembolsoJurosCP                string `bson:"qtd_reembolso_juros_cp"`                 // Quantidade de reembolso de juros CP
	QtdReembolsoJurosCPAgrupado        string `bson:"qtd_reembolso_juros_cp_agrupado"`        // Quantidade de reembolso de juros CP agrupado
	QtdReembolsoRendimentos            string `bson:"qtd_reembolso_rendimentos"`              // Quantidade de reembolso de rendimentos
	QtdReembolsoRendimentosAgrupado    string `bson:"qtd_reembolso_rendimentos_agrupado"`     // Quantidade de reembolso de rendimentos agrupado
	QtdResgateIntegralizacao           string `bson:"qtd_resgate_integralizacao"`             // Quantidade de resgate para integralização
	QtdResgateParcial                  string `bson:"qtd_resgate_parcial"`                    // Quantidade de resgate parcial
	QtdRevolucaoCotst                  string `bson:"qtd_revolucao_cotst"`                    // Quantidade de revolução de cotistas
	QtdRendimentosAuferidos            string `bson:"qtd_rendimentos_auferidos"`              // Quantidade de rendimentos auferidos
	QtdResgates                        string `bson:"qtd_resgates"`                           // Quantidade de resgates
	QtdSuspensaoCota                   string `bson:"qtd_suspensao_cota"`                     // Quantidade de suspensão de cota
	QtdVendaCota                       string `bson:"qtd_venda_cota"`                         // Quantidade de venda de cota
	QtdAmortizacaoPrecoAdquirido       string `bson:"qtd_amortizacao_preco_adquirido"`        // Quantidade de amortização de preço adquirido
	QtdReinvestimento                  string `bson:"qtd_reinvestimento"`                     // Quantidade de reinvestimento
	QtdSubscricao                      string `bson:"qtd_subscricao"`                         // Quantidade de subscrição
	TipoComitente1                     string `bson:"tipo_comitente_1"`                       // Tipo de comitente 1
	TipoComitente2                     string `bson:"tipo_comitente_2"`                       // Tipo de comitente 2
	TipoComitente3                     string `bson:"tipo_comitente_3"`                       // Tipo de comitente 3
	TipoControleGestao                 string `bson:"tipo_controle_gestao"`                   // Tipo de controle de gestão
	TipoFundo                          string `bson:"tipo_fundo"`                             // Tipo de fundo
	TipoRentabFund                     string `bson:"tipo_rentab_fund"`                       // Tipo de rentabilidade do fundo
	VencimentoCotas                    string `bson:"vencimento_cotas"`                       // Vencimento das cotas
}
