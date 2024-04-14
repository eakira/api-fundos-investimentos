package entity

type LaminaEntity struct {
	AnoAnterExemplo           string `bson:"ano_anterior_exemplo,omitempty"`
	AnoExemplo                string `bson:"ano_exemplo,omitempty"`
	AnoSemRentab              string `bson:"ano_sem_rentab,omitempty"`
	CalcRentabFundo           string `bson:"calc_rentab_fundo,omitempty"`
	CalcRentabFundoGatilho    string `bson:"calc_rentab_fundo_gatilho,omitempty"`
	ClasseRiscoAdmin          string `bson:"classe_risco_admin,omitempty"`
	FundoCnpj                 string `bson:"cnpj_fundo,omitempty"`
	CondicCaren               string `bson:"condic_caren,omitempty"`
	CondicEntr                string `bson:"condic_entr,omitempty"`
	CondicSaida               string `bson:"condic_saida,omitempty"`
	ConflitoVenda             string `bson:"conflito_venda,omitempty"`
	ConversaoCotaCanc         string `bson:"conversao_cota_canc,omitempty"`
	ConversaoCotaCompra       string `bson:"conversao_cota_compra,omitempty"`
	DenomSocial               string `bson:"denom_social,omitempty"`
	DerivProtecaoCarteira     string `bson:"deriv_protecao_carteira,omitempty"`
	DistribGestorUnico        string `bson:"distrib_gestor_unico,omitempty"`
	DsRentabGatilho           string `bson:"ds_rentab_gatilho,omitempty"`
	DtComptc                  string `bson:"dt_comptc,omitempty"`
	DtFimDespesa              string `bson:"dt_fim_despesa,omitempty"`
	DtIniAtiv5Ano             string `bson:"dt_ini_ativ_5_ano,omitempty"`
	DtIniDespesa              string `bson:"dt_ini_despesa,omitempty"`
	EnderEletronico           string `bson:"ender_eletronico,omitempty"`
	EnderEletronicoDespesa    string `bson:"ender_eletronico_despesa,omitempty"`
	EnderEletronicoReclamacao string `bson:"ender_eletronico_reclamacao,omitempty"`
	HoraAplicResgate          string `bson:"hora_aplic_resgate,omitempty"`
	IndiceRefer               string `bson:"indice_refer,omitempty"`
	InvestAdic                string `bson:"invest_adic,omitempty"`
	InvestInicialMin          string `bson:"invest_inicial_min,omitempty"`
	NmFantasia                string `bson:"nm_fantasia,omitempty"`
	Objetivo                  string `bson:"objetivo,omitempty"`
	PolitInvest               string `bson:"polit_invest,omitempty"`
	PrAtivoEmissor            string `bson:"pr_ativo_emissor,omitempty"`
	PrPlAlavanc               string `bson:"pr_pl_alavanc,omitempty"`
	PrPlAplicMaxFundoUnico    string `bson:"pr_pl_aplic_max_fundo_unico,omitempty"`
	PrPlAtivoCredPriv         string `bson:"pr_pl_ativo_cred_priv,omitempty"`
	PrPlAtivoExterior         string `bson:"pr_pl_ativo_exterior,omitempty"`
	PrPlDespesa               string `bson:"pr_pl_despesa,omitempty"`
	PrRentabFundo5Ano         string `bson:"pr_rentab_fundo_5_ano,omitempty"`
	PrVariacaoIndiceRefer5Ano string `bson:"pr_variacao_indice_refer_5_ano,omitempty"`
	PrVariacaoPerfm           string `bson:"pr_variacao_perfm,omitempty"`
	PublicoAlvo               string `bson:"publico_alvo,omitempty"`
	QtAnoPerda                string `bson:"qt_ano_perda,omitempty"`
	QtDiaCaren                string `bson:"qt_dia_caren,omitempty"`
	QtDiaConversaoCotaCompra  string `bson:"qt_dia_conversao_cota_compra,omitempty"`
	QtDiaConversaoCotaResgate string `bson:"qt_dia_conversao_cota_resgate,omitempty"`
	QtDiaPagtoResgate         string `bson:"qt_dia_pagto_resgate,omitempty"`
	QtDiaSaida                string `bson:"qt_dia_saida,omitempty"`
	RemunDistrib              string `bson:"remun_distrib,omitempty"`
	RentabGatilho             string `bson:"rentab_gatilho,omitempty"`
	ResgateMin                string `bson:"resgate_min,omitempty"`
	RestrInvest               string `bson:"restr_invest,omitempty"`
	RiscoPerda                string `bson:"risco_perda,omitempty"`
	RiscoPerdaNegativo        string `bson:"risco_perda_negativo,omitempty"`
	TaxaAdm                   string `bson:"taxa_adm,omitempty"`
	TaxaAdmMax                string `bson:"taxa_adm_max,omitempty"`
	TaxaAdmMin                string `bson:"taxa_adm_min,omitempty"`
	TaxaAdmObs                string `bson:"taxa_adm_obs,omitempty"`
	TaxaEntr                  string `bson:"taxa_entr,omitempty"`
	TaxaPerfm                 string `bson:"taxa_perfm,omitempty"`
	TaxaSaida                 string `bson:"taxa_saida,omitempty"`
	TelSac                    string `bson:"tel_sac,omitempty"`
	TpDiaPagtoResgate         string `bson:"tp_dia_pagto_resgate,omitempty"`
	TpTaxaAdm                 string `bson:"tp_taxa_adm,omitempty"`
	VlAjustePerfmExemplo      string `bson:"vl_ajuste_perfm_exemplo,omitempty"`
	VlDespesa3Ano             string `bson:"vl_despesa_3_ano,omitempty"`
	VlDespesa5Ano             string `bson:"vl_despesa_5_ano,omitempty"`
	VlDespesaExemplo          string `bson:"vl_despesa_exemplo,omitempty"`
	VlImpostoExemplo          string `bson:"vl_imposto_exemplo,omitempty"`
	VlMinPerman               string `bson:"vl_min_perman,omitempty"`
	VlPatrimLiq               string `bson:"vl_patrim_liq,omitempty"`
	VlResgateExemplo          string `bson:"vl_resgate_exemplo,omitempty"`
	VlRetorno3Ano             string `bson:"vl_retorno_3_ano,omitempty"`
	VlRetorno5Ano             string `bson:"vl_retorno_5_ano,omitempty"`
	VlTaxaEntrExemplo         string `bson:"vl_taxa_entr_exemplo,omitempty"`
	VlTaxaSaidaExemplo        string `bson:"vl_taxa_saida_exemplo,omitempty"`
	InfSac                    string `bson:"inf_sac,omitempty"`
}

type LaminaCarteiraEntity struct {
	FundoCnpj   string `bson:"cnpj_fundo,omitempty"`
	DenomSocial string `bson:"denom_social,omitempty"`
	DtComptc    string `bson:"dt_comptc,omitempty"`
	PrPlAtivo   string `bson:"pr_pl_ativo,omitempty"`
	TpAtivo     string `bson:"tp_ativo,omitempty"`
}

type LaminaRentabilidadeAnoEntity struct {
	AnoRentab                string `bson:"ano_rentab,omitempty"`
	FundoCnpj                string `bson:"cnpj_fundo,omitempty"`
	DenomSocial              string `bson:"denom_social,omitempty"`
	DtComptc                 string `bson:"dt_comptc,omitempty"`
	PrPerfmIndiceReferAno    string `bson:"pr_perfm_indice_refer_ano,omitempty"`
	PrRentabAno              string `bson:"pr_rentab_ano,omitempty"`
	PrVariacaoIndiceReferAno string `bson:"pr_variacao_indice_refer_ano,omitempty"`
	RentabAnoObs             string `bson:"rentab_ano_obs,omitempty"`
}

type LaminaRentabilidadeMesEntity struct {
	FundoCnpj                string `bson:"cnpj_fundo,omitempty"`
	DenomSocial              string `bson:"denom_social,omitempty"`
	DtComptc                 string `bson:"dt_comptc,omitempty"`
	MesRentab                string `bson:"mes_rentab,omitempty"`
	PrPerfmIndiceReferMes    string `bson:"pr_perfm_indice_refer_mes,omitempty"`
	PrRentabMes              string `bson:"pr_rentab_mes,omitempty"`
	PrVariacaoIndiceReferMes string `bson:"pr_variacao_indice_refer_mes,omitempty"`
	RentabMesObs             string `bson:"rentab_mes_obs,omitempty"`
}
