package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InformacaoDiariaEntity struct {
	ID                     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FundoCnpj              string             `bson:"fundo_cnpj,omitempty"`
	CaptacaoDia            float64            `bson:"captacao_dia,omitempty"`
	DataCompetencia        time.Time          `bson:"data_competencia,omitempty"`
	ResgateDia             float64            `bson:"resgate_dia,omitempty"`
	TipoFundo              string             `bson:"tipo_fundo,omitempty"`
	ValorPatrimonioLiquido float64            `bson:"valor_patrimonio_liquido,omitempty"`
	CotaValor              float64            `bson:"cota_valor,omitempty"`
	QtdCotista             int                `bson:"qtd_cotista,omitempty"`
	ValorTotal             float64            `bson:"valor_total,omitempty"`
	CreatedAt              time.Time          `bson:"created_at,omitempty"`
	UpdatedAt              time.Time          `bson:"update_at,omitempty"`
}
