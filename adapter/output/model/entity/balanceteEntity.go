package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type BalacenteEntity struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CodigoConta     string             `bson:"codigo_conta"`
	FundoCnpj       string             `bson:"fundo_cnpj"`
	DataCompetencia string             `bson:"data_competencia"`
	PlanoContabil   string             `bson:"plano_contabil"`
	SaldoConta      string             `bson:"saldo_conta"`
}
