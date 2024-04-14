package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BalacenteEntity struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CodigoConta     int                `bson:"codigo_conta"`
	FundoCnpj       string             `bson:"fundo_cnpj"`
	DataCompetencia time.Time          `bson:"data_competencia"`
	PlanoContabil   string             `bson:"plano_contabil"`
	SaldoConta      float64            `bson:"saldo_conta"`
}
