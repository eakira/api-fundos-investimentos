package domain

import "time"

type ExtratoDomain struct {
	Id                             string
	LimMaxCreditoPrivado           float64
	LimMaxAtivosExterior           float64
	LimMaxMesmoGestor              float64
	AplicacaoMinima                float64
	PodeCreditoprivado             string
	CotaDiaria                     string
	MetodoCalculoPerformance       string
	ClasseAnbima                   string
	FundoCnpj                      string
	Codominio                      string
	ContraparteLigado              string
	CotaValorEmissao               string
	CotaPatriminioLiquido          string
	DenominacaoSocial              string
	FormaDistribuicao              string
	DataCompetenciaDocumento       time.Time
	TaxaDeIngresso                 string
	TaxaDePerformance              string
	TaxaDeSaida                    string
	FatorLimiteOperacoes           float64
	FinalidadeOperacoesDerivativos string
	FundoCotas                     string
	FundoEspelho                   string
	InformacoesAdicionais          string
	PodeInvestirExterior           string
	NomeDoMercado                  string
	NegociacoesMercadoOrganizado   string
	PodeDerivativos                string
	PodeOperacoesSuperiorPl        string
	ParametroTaxaPerformance       string
	PoliticaInvistimento           string
	PrAcaoMax                      float64
	PrAcaoMin                      float64
	PrAdminGestorMax               float64
	PrAdminGestorMin               float64
	PrAtivoOutroMax                float64
	PrAtivoOutroMin                float64
	PrCiaMax                       float64
	PrCiaMin                       float64
	PrCompromMax                   float64
	PrCompromMin                   float64
	PrCotaEtfMax                   float64
	PrCotaEtfMin                   float64
	PrCotaFiMax                    float64
	PrCotaFiMin                    float64
	PrCotaFiProfMax                float64
	PrCotaFiProfMin                float64
	PrCotaFiQualifMax              float64
	PrCotaFiQualifMin              float64
	PrCotaFicMax                   float64
	PrCotaFicMin                   float64
	PrCotaFicProfMax               float64
	PrCotaFicProfMin               float64
	PrCotaFicQualifMax             float64
	PrCotaFicQualifMin             float64
	PrCotaFicfidcMax               float64
	PrCotaFicfidcMin               float64
	PrCotaFicfidcNpMax             float64
	PrCotaFicfidcNpMin             float64
	PrCotaFicfipMax                float64
	PrCotaFicfipMin                float64
	PrCotaFidcMax                  float64
	PrCotaFidcMin                  float64
	PrCotaFidcNpMax                float64
	PrCotaFidcNpMin                float64
	PrCotaFiiMax                   float64
	PrCotaFiiMin                   float64
	PrCotaFipMax                   float64
	PrCotaFipMin                   float64
	PrCotaFmieeMax                 float64
	PrCotaFmieeMin                 float64
	PrCriMax                       float64
	PrCriMin                       float64
	PrDebentureMax                 float64
	PrDebentureMin                 float64
	PrDerivMax                     float64
	PrDerivMin                     float64
	PrEmissorOutroMax              float64
	PrEmissorOutroMin              float64
	PrFiMax                        float64
	PrFiMin                        float64
	PrIndiceReferTaxaPerfm         float64
	PrInstituicaoFinancMax         float64
	PrInstituicaoFinancMin         float64
	PrNpMax                        float64
	PrNpMin                        float64
	PrOuroMax                      float64
	PrOuroMin                      float64
	PrTitInstituicaoFinancBacenMax float64
	PrTitInstituicaoFinancBacenMin float64
	PrTitpubMax                    float64
	PrTitpubMin                    float64
	PrUniaoMax                     float64
	PrUniaoMin                     float64
	PrVlmobMax                     float64
	PrVlmobMin                     float64
	Prazo                          string
	PrazoAtualizCota               string
	PublicoAlvo                    string
	QtDiaConversaoCota             int
	PrUniaoMaxQtDiaPagtoCota       int
	PrUniaoMaxQtDiaPagtoResgate    int
	PrUniaoMaxQtDiaResgateCotas    int
	RegAnbima                      string
	ResultCartIncorpPl             string
	TaxaAdm                        float64
	TaxaCustodiaMax                float64
	TaxaIngressoPr                 float64
	TaxaIngressoReal               float64
	TaxaPerfm                      float64
	TaxaSaidaPagtoResgate          string
	TaxaSaidaPr                    float64
	TaxaSaidaReal                  float64
	TpDiaPagtoResgate              string
	TpPrazo                        string
	VlCupom                        float64
}
