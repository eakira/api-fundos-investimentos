package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FundosEntity struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	AdministradorNome     string             `bson:"administrador_nome"`
	AuditorNome           string             `bson:"auditor_nome"`
	CodigoCVM             int                `bson:"codigo_cvm"`
	Classe                string             `bson:"classe"`
	ClasseAnbima          string             `bson:"classe_anbima"`
	AdministradorCNPJ     string             `bson:"administrador_cnpj"`
	AuditorCNPJ           string             `bson:"auditor_cnpj"`
	ControladorCNPJ       string             `bson:"controlador_cnpj"`
	CustodianteCNPJ       string             `bson:"custodiante_cnpj"`
	FundoCNPJ             string             `bson:"fundo_cnpj"`
	Condominio            string             `bson:"condominio"`
	Controlador           string             `bson:"controlador"`
	GestorCPFCNPJ         string             `bson:"gestor_cpf_cnpj"`
	Custodiante           string             `bson:"custodiante"`
	FundoNome             string             `bson:"fundo_nome"`
	Diretor               string             `bson:"diretor"`
	DataCancelamento      time.Time          `bson:"data_cancelamento"`
	DataConstituicao      time.Time          `bson:"data_constituicao"`
	DataFim               time.Time          `bson:"data_fim"`
	DataInicio            time.Time          `bson:"data_inicio"`
	DataInicioClasse      time.Time          `bson:"data_inicio_classe"`
	DataInicioSocial      time.Time          `bson:"data_inicio_social"`
	DataInicioSituacao    time.Time          `bson:"data_inicio_situacao"`
	DataPatrimonioLiquido time.Time          `bson:"data_patrimonio_liquido"`
	DataRegistro          time.Time          `bson:"data_registro"`
	Entidade              string             `bson:"entidade"`
	CotasPossui           string             `bson:"cotas_possui"`
	Exclusivo             string             `bson:"exclusivo"`
	GestorNome            string             `bson:"gestor_nome"`
	InformacoesAdicionais string             `bson:"informacoes_adicionais"`
	InfoTaxaPerformance   string             `bson:"info_taxa_performance"`
	Exterior100           string             `bson:"exterior_100"`
	PessoaFouPJ           string             `bson:"pf_pj"`
	PublicoAlvo           string             `bson:"publico_alvo"`
	IndicadorDesempenho   string             `bson:"indicador_desempenho"`
	Situacao              string             `bson:"situacao"`
	TaxaAdm               string             `bson:"taxa_adm"`
	TaxaPerformance       string             `bson:"taxa_performance"`
	TipoFundo             string             `bson:"tipo_fundo"`
	TributacaoLongoPrazo  string             `bson:"tributacao_longo_prazo"`
	ValorPatriminioLiq    float64            `bson:"valor_patriminio_liquido"`
}
