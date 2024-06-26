package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArquivosEntity struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Endereco    string             `bson:"endereco"`
	TipoArquivo string             `bson:"tipo-arquivo"`
	Referencia  string             `bson:"referencia"`
	Status      string             `bson:"status"`
	Baixar      bool               `bson:"baixar"`
	Download    bool               `bson:"download"`
	Processado  bool               `bson:"processado"`
	CreatedAt   time.Time          `bson:"created-at"`
	UpdatedAt   time.Time          `bson:"update-at"`
}
