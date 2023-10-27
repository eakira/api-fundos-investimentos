package entity

type BalacenteEntity struct {
	CodigoConta     string `bson:"codigo_conta"`
	FundoCnpj       string `bson:"fundo_cnpj"`
	DataCompetencia string `bson:"data_competencia"`
	PlanoContabil   string `bson:"plano_contabil"`
	SaldoConta      string `bson:"saldo_conta"`
}
