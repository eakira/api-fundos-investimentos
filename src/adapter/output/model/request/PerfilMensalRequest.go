package request

type PerfilMensalRequest struct {
	CenarioFPRCupom                    string `json:"cenario_fpr_cupom"`
	CenarioFPRDolar                    string `json:"cenario_fpr_dolar"`
	CenarioFPRIBovespa                 string `json:"cenario_fpr_ibovespa"`
	CenarioFPRJuros                    string `json:"cenario_fpr_juros"`
	CenarioFPROutro                    string `json:"cenario_fpr_outro"`
	CNPJFundo                          string `json:"cnpj_fundo"`
	ComitenteLigado1                   string `json:"comitente_ligado_1"`
	ComitenteLigado2                   string `json:"comitente_ligado_2"`
	ComitenteLigado3                   string `json:"comitente_ligado_3"`
	CPFCNPJComitente1                  string `json:"cpf_cnpj_comitente_1"`
	CPFCNPJComitente2                  string `json:"cpf_cnpj_comitente_2"`
	CPFCNPJComitente3                  string `json:"cpf_cnpj_comitente_3"`
	CPFCNPJEmissor1                    string `json:"cpf_cnpj_emissor_1"`
	CPFCNPJEmissor2                    string `json:"cpf_cnpj_emissor_2"`
	CPFCNPJEmissor3                    string `json:"cpf_cnpj_emissor_3"`
	DelibAssemb                        string `json:"delib_assemb"`
	DenomSocial                        string `json:"denom_social"`
	DataComptc                         string `json:"dt_comptc"`
	DataCotaTaxaPerfm                  string `json:"dt_cota_taxa_perfm"`
	EmissorLigado1                     string `json:"emissor_ligado_1"`
	EmissorLigado2                     string `json:"emissor_ligado_2"`
	EmissorLigado3                     string `json:"emissor_ligado_3"`
	FatorRiscoNocional                 string `json:"fator_risco_nocional"`
	FatorRiscoOutro                    string `json:"fator_risco_outro"`
	FPR                                string `json:"fpr"`
	JustifVotoAdminAssemb              string `json:"justif_voto_admin_assemb"`
	ModVar                             string `json:"mod_var"`
	NrCotstBanco                       string `json:"nr_cotst_banco"`
	NrCotstCapitaliz                   string `json:"nr_cotst_capitaliz"`
	NrCotstCorretoraDistrib            string `json:"nr_cotst_corretora_distrib"`
	NrCotstDistrib                     string `json:"nr_cotst_distrib"`
	NrCotstEAPC                        string `json:"nr_cotst_eapc"`
	NrCotstEFPC                        string `json:"nr_cotst_efpc"`
	NrCotstEntidPrevidCompl            string `json:"nr_cotst_entid_previd_compl"`
	NrCotstFIClube                     string `json:"nr_cotst_fi_clube"`
	NrCotstInvnr                       string `json:"nr_cotst_invnr"`
	NrCotstOutro                       string `json:"nr_cotst_outro"`
	NrCotstPFPB                        string `json:"nr_cotst_pf_pb"`
	NrCotstPFVarejo                    string `json:"nr_cotst_pf_varejo"`
	NrCotstPJFinanc                    string `json:"nr_cotst_pj_financ"`
	NrCotstPJNaoFinancPB               string `json:"nr_cotst_pj_nao_financ_pb"`
	NrCotstPJNaoFinancVarejo           string `json:"nr_cotst_pj_nao_financ_varejo"`
	NrCotstRPPS                        string `json:"nr_cotst_rpps"`
	NrCotstSegur                       string `json:"nr_cotst_segur"`
	NrDiaCemPerc                       string `json:"nr_dia_cem_perc"`
	NrDiaCinquPerc                     string `json:"nr_dia_cinqu_perc"`
	PFPJComitente1                     string `json:"pf_pj_comitente_1"`
	PFPJComitente2                     string `json:"pf_pj_comitente_2"`
	PFPJComitente3                     string `json:"pf_pj_comitente_3"`
	QtdAcoesEmitidasMaiorQtdAutoriz    string `json:"qtd_acoes_emitidas_maior_qtd_autoriz"`
	QtdCotstEmit                       string `json:"qtd_cotst_emit"`
	QtdCotstResgatadas                 string `json:"qtd_cotst_resgatadas"`
	QtdDiasCorridos                    string `json:"qtd_dias_corridos"`
	QtdEmissaoResgate                  string `json:"qtd_emissao_resgate"`
	QtdEmissaoResgateAgrupado          string `json:"qtd_emissao_resgate_agrupado"`
	QtdRecolhimentoIntegralizacao      string `json:"qtd_recolhimento_integralizacao"`
	QtdReembolsoDividendosJSCP         string `json:"qtd_reembolso_dividendos_jscp"`
	QtdReembolsoDividendosJSCPAgrupado string `json:"qtd_reembolso_dividendos_jscp_agrupado"`
	QtdReembolsoJurosCP                string `json:"qtd_reembolso_juros_cp"`
	QtdReembolsoJurosCPAgrupado        string `json:"qtd_reembolso_juros_cp_agrupado"`
	QtdReembolsoRendimentos            string `json:"qtd_reembolso_rendimentos"`
	QtdReembolsoRendimentosAgrupado    string `json:"qtd_reembolso_rendimentos_agrupado"`
	QtdResgateIntegralizacao           string `json:"qtd_resgate_integralizacao"`
	QtdResgateParcial                  string `json:"qtd_resgate_parcial"`
	QtdRevolucaoCotst                  string `json:"qtd_revolucao_cotst"`
	QtdRendimentosAuferidos            string `json:"qtd_rendimentos_auferidos"`
	QtdResgates                        string `json:"qtd_resgates"`
	QtdSuspensaoCota                   string `json:"qtd_suspensao_cota"`
	QtdVendaCota                       string `json:"qtd_venda_cota"`
	QtdAmortizacaoPrecoAdquirido       string `json:"qtd_amortizacao_preco_adquirido"`
	QtdReinvestimento                  string `json:"qtd_reinvestimento"`
	QtdSubscricao                      string `json:"qtd_subscricao"`
	TipoComitente1                     string `json:"tipo_comitente_1"`
	TipoComitente2                     string `json:"tipo_comitente_2"`
	TipoComitente3                     string `json:"tipo_comitente_3"`
	TipoControleGestao                 string `json:"tipo_controle_gestao"`
	TipoFundo                          string `json:"tipo_fundo"`
	TipoRentabFund                     string `json:"tipo_rentab_fund"`
	VencimentoCotas                    string `json:"vencimento_cotas"`
}
