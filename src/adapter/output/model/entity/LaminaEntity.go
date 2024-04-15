package entity

import "time"

type LaminaEntity struct {
	AnoAnterExemplo           int       `bson:"ano_anterior_exemplo,omitempty"`
	AnoExemplo                int       `bson:"ano_exemplo,omitempty"`
	AnoSemRentab              string    `bson:"ano_sem_rentab,omitempty"`
	CalcRentabFundo           string    `bson:"calc_rentab_fundo,omitempty"`
	CalcRentabFundoGatilho    string    `bson:"calc_rentab_fundo_gatilho,omitempty"`
	ClasseRiscoAdmin          int       `bson:"classe_risco_admin,omitempty"`
	FundoCnpj                 string    `bson:"cnpj_fundo,omitempty"`
	CondicCaren               string    `bson:"condic_caren,omitempty"`
	CondicEntr                string    `bson:"condic_entr,omitempty"`
	CondicSaida               string    `bson:"condic_saida,omitempty"`
	ConflitoVenda             string    `bson:"conflito_venda,omitempty"`
	ConversaoCotaCanc         string    `bson:"conversao_cota_canc,omitempty"`
	ConversaoCotaCompra       string    `bson:"conversao_cota_compra,omitempty"`
	DenomSocial               string    `bson:"denom_social,omitempty"`
	DerivProtecaoCarteira     string    `bson:"deriv_protecao_carteira,omitempty"`
	DistribGestorUnico        string    `bson:"distrib_gestor_unico,omitempty"`
	DsRentabGatilho           string    `bson:"ds_rentab_gatilho,omitempty"`
	DtComptc                  time.Time `bson:"dt_comptc,omitempty"`
	DtFimDespesa              time.Time `bson:"dt_fim_despesa,omitempty"`
	DtIniAtiv5Ano             time.Time `bson:"dt_ini_ativ_5_ano,omitempty"`
	DtIniDespesa              time.Time `bson:"dt_ini_despesa,omitempty"`
	EnderEletronico           string    `bson:"ender_eletronico,omitempty"`
	EnderEletronicoDespesa    string    `bson:"ender_eletronico_despesa,omitempty"`
	EnderEletronicoReclamacao string    `bson:"ender_eletronico_reclamacao,omitempty"`
	HoraAplicResgate          string    `bson:"hora_aplic_resgate,omitempty"`
	IndiceRefer               string    `bson:"indice_refer,omitempty"`
	InvestAdic                float64   `bson:"invest_adic,omitempty"`
	InvestInicialMin          float64   `bson:"invest_inicial_min,omitempty"`
	NmFantasia                string    `bson:"nm_fantasia,omitempty"`
	Objetivo                  string    `bson:"objetivo,omitempty"`
	PolitInvest               string    `bson:"polit_invest,omitempty"`
	PrAtivoEmissor            float64   `bson:"pr_ativo_emissor,omitempty"`
	PrPlAlavanc               float64   `bson:"pr_pl_alavanc,omitempty"`
	PrPlAplicMaxFundoUnico    float64   `bson:"pr_pl_aplic_max_fundo_unico,omitempty"`
	PrPlAtivoCredPriv         float64   `bson:"pr_pl_ativo_cred_priv,omitempty"`
	PrPlAtivoExterior         float64   `bson:"pr_pl_ativo_exterior,omitempty"`
	PrPlDespesa               float64   `bson:"pr_pl_despesa,omitempty"`
	PrRentabFundo5Ano         float64   `bson:"pr_rentab_fundo_5_ano,omitempty"`
	PrVariacaoIndiceRefer5Ano float64   `bson:"pr_variacao_indice_refer_5_ano,omitempty"`
	PrVariacaoPerfm           float64   `bson:"pr_variacao_perfm,omitempty"`
	PublicoAlvo               string    `bson:"publico_alvo,omitempty"`
	QtAnoPerda                int       `bson:"qt_ano_perda,omitempty"`
	QtDiaCaren                int       `bson:"qt_dia_caren,omitempty"`
	QtDiaConversaoCotaCompra  int       `bson:"qt_dia_conversao_cota_compra,omitempty"`
	QtDiaConversaoCotaResgate int       `bson:"qt_dia_conversao_cota_resgate,omitempty"`
	QtDiaPagtoResgate         int       `bson:"qt_dia_pagto_resgate,omitempty"`
	QtDiaSaida                int       `bson:"qt_dia_saida,omitempty"`
	RemunDistrib              string    `bson:"remun_distrib,omitempty"`
	RentabGatilho             string    `bson:"rentab_gatilho,omitempty"`
	ResgateMin                float64   `bson:"resgate_min,omitempty"`
	RestrInvest               string    `bson:"restr_invest,omitempty"`
	RiscoPerda                string    `bson:"risco_perda,omitempty"`
	RiscoPerdaNegativo        string    `bson:"risco_perda_negativo,omitempty"`
	TaxaAdm                   float64   `bson:"taxa_adm,omitempty"`
	TaxaAdmMax                float64   `bson:"taxa_adm_max,omitempty"`
	TaxaAdmMin                float64   `bson:"taxa_adm_min,omitempty"`
	TaxaAdmObs                float64   `bson:"taxa_adm_obs,omitempty"`
	TaxaEntr                  float64   `bson:"taxa_entr,omitempty"`
	TaxaPerfm                 float64   `bson:"taxa_perfm,omitempty"`
	TaxaSaida                 float64   `bson:"taxa_saida,omitempty"`
	TelSac                    string    `bson:"tel_sac,omitempty"`
	TpDiaPagtoResgate         string    `bson:"tp_dia_pagto_resgate,omitempty"`
	TpTaxaAdm                 string    `bson:"tp_taxa_adm,omitempty"`
	VlAjustePerfmExemplo      float64   `bson:"vl_ajuste_perfm_exemplo,omitempty"`
	VlDespesa3Ano             float64   `bson:"vl_despesa_3_ano,omitempty"`
	VlDespesa5Ano             float64   `bson:"vl_despesa_5_ano,omitempty"`
	VlDespesaExemplo          float64   `bson:"vl_despesa_exemplo,omitempty"`
	VlImpostoExemplo          float64   `bson:"vl_imposto_exemplo,omitempty"`
	VlMinPerman               float64   `bson:"vl_min_perman,omitempty"`
	VlPatrimLiq               float64   `bson:"vl_patrim_liq,omitempty"`
	VlResgateExemplo          float64   `bson:"vl_resgate_exemplo,omitempty"`
	VlRetorno3Ano             float64   `bson:"vl_retorno_3_ano,omitempty"`
	VlRetorno5Ano             float64   `bson:"vl_retorno_5_ano,omitempty"`
	VlTaxaEntrExemplo         float64   `bson:"vl_taxa_entr_exemplo,omitempty"`
	VlTaxaSaidaExemplo        float64   `bson:"vl_taxa_saida_exemplo,omitempty"`
	InfSac                    string    `bson:"inf_sac,omitempty"`
}

type LaminaCarteiraEntity struct {
	FundoCnpj   string    `bson:"cnpj_fundo,omitempty"`
	DenomSocial string    `bson:"denom_social,omitempty"`
	DtComptc    time.Time `bson:"dt_comptc,omitempty"`
	PrPlAtivo   float64   `bson:"pr_pl_ativo,omitempty"`
	TpAtivo     string    `bson:"tp_ativo,omitempty"`
}

type LaminaRentabilidadeAnoEntity struct {
	AnoRentab                int       `bson:"ano_rentab,omitempty"`
	FundoCnpj                string    `bson:"cnpj_fundo,omitempty"`
	DenomSocial              string    `bson:"denom_social,omitempty"`
	DtComptc                 time.Time `bson:"dt_comptc,omitempty"`
	PrPerfIndiceReferAno     float64   `bson:"pr_perfm_indice_refer_ano,omitempty"`
	PrRentabAno              float64   `bson:"pr_rentab_ano,omitempty"`
	PrVariacaoIndiceReferAno float64   `bson:"pr_variacao_indice_refer_ano,omitempty"`
	RentabAnoObservacoes     string    `bson:"rentab_ano_obs,omitempty"`
}

type LaminaRentabilidadeMesEntity struct {
	FundoCnpj                string    `bson:"cnpj_fundo,omitempty"`
	DenomSocial              string    `bson:"denom_social,omitempty"`
	DtComptc                 time.Time `bson:"dt_comptc,omitempty"`
	MesReferencia            int       `bson:"mes_rentab,omitempty"`
	PrPerfIndiceReferMes     float64   `bson:"pr_perfm_indice_refer_mes,omitempty"`
	PrRentabMes              float64   `bson:"pr_rentab_mes,omitempty"`
	PrVariacaoIndiceReferMes float64   `bson:"pr_variacao_indice_refer_mes,omitempty"`
	RentabMesObservacoes     string    `bson:"rentab_mes_obs,omitempty"`
}
