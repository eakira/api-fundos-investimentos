package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InformacoesFundoEntity struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty"`
	AgenciaRisco               string             `bson:"agencia_risco,omitempty"`
	ApresentacaoAdministrador  string             `bson:"apresentacao_admin,omitempty"`
	ApresentacaoGestor         string             `bson:"apresentacao_gestor,omitempty"`
	CnpjAgenciaRisco           string             `bson:"cnpj_agencia_risco,omitempty"`
	FundoCnpj                  string             `bson:"cnpj_fundo,omitempty"`
	DisclaimerAgenciaRisco     string             `bson:"disclaimer_agencia_risco,omitempty"`
	DistribuidorLigado         string             `bson:"distribuidor_ligado,omitempty"`
	DataCompetencia            time.Time          `bson:"data_competencia,omitempty"`
	GrauRisco                  string             `bson:"grau_risco,omitempty"`
	InformacoesAutorregulacao  string             `bson:"informacoes_autorregulacao,omitempty"`
	OutrasInformacoes          string             `bson:"outras_informacoes,omitempty"`
	PeriodoMinimoDivulgacao    string             `bson:"periodo_min_divulgacao,omitempty"`
	PoliticaAdministracaoRisco string             `bson:"politica_admin_risco,omitempty"`
	PoliticaDistribuicaoCotas  string             `bson:"politica_distribuicao_cotas,omitempty"`
	PoliticaVoto               string             `bson:"politica_voto,omitempty"`
	RiscoCarteira              string             `bson:"risco_carteira,omitempty"`
	ExistenciaAgenciaRisco     string             `bson:"existencia_agencia_risco,omitempty"`
	TipoFundo                  string             `bson:"tipo_fundo,omitempty"`
	TributacaoFundoCotista     string             `bson:"tributacao_fundo_cotista,omitempty"`
	VotoGestor                 string             `bson:"voto_gestor,omitempty"`
}

type InformacoesDivulgacaoEntity struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	FundoCnpj        string             `bson:"cnpj_fundo,omitempty"`
	DataCompetencia  time.Time          `bson:"data_competencia,omitempty"`
	FormaInformacoes string             `bson:"forma_informacoes,omitempty"`
	LocalInformacoes string             `bson:"local_informacoes,omitempty"`
	MeioInformacoes  string             `bson:"meio_informacoes,omitempty"`
	TipoFundo        string             `bson:"tipo_fundo,omitempty"`
}

type InformacoesCotistaEntity struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	FundoCnpj            string             `bson:"cnpj_fundo,omitempty"`
	DescricaoRespCotista string             `bson:"descricao_resp_cotista,omitempty"`
	DataCompetencia      time.Time          `bson:"data_competencia,omitempty"`
	FormaInfCotista      string             `bson:"forma_inf_cotista,omitempty"`
	LocalInfCotista      string             `bson:"local_inf_cotista,omitempty"`
	MeioInfCotista       string             `bson:"meio_inf_cotista,omitempty"`
	TipoFundo            string             `bson:"tipo_fundo,omitempty"`
}
type ServicoPrestadoEntity struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	FundoCnpj        string             `bson:"cnpj_fundo,omitempty"`
	DescricaoServico string             `bson:"descricao_servico,omitempty"`
	DataCompetencia  time.Time          `bson:"data_competencia,omitempty"`
	NomePrestador    string             `bson:"nome_prestador,omitempty"`
	TipoFundo        string             `bson:"tipo_fundo,omitempty"`
}
