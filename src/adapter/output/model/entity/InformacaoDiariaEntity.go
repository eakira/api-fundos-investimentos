package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type InformacaoDiariaEntity struct {
	ID                     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FundoCnpj              string             `bson:"fundo_cnpj,omitempty"`
	CaptacaoDia            string             `bson:"captacao_dia,omitempty"`
	DataCompetencia        string             `bson:"data_competencia,omitempty"`
	ResgateDia             string             `bson:"resgate_dia,omitempty"`
	TipoFundo              string             `bson:"tipo_fundo,omitempty"`
	ValorPatrimonioLiquido string             `bson:"valor_patrimonio_liquido,omitempty"`
	CotaValor              string             `bson:"cota_valor,omitempty"`
	QtdCotista             string             `bson:"qtd_cotista,omitempty"`
	ValorTotal             string             `bson:"valor_total,omitempty"`
}
