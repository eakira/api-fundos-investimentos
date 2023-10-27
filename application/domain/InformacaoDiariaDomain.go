package domain

import "time"

type InformacaoDiariaDomain struct {
	Id                     string
	FundoCnpj              string
	CaptacaoDia            string
	DataCompetencia        string
	ResgateDia             string
	TipoFundo              string
	ValorPatrimonioLiquido string
	CotaValor              string
	QtdCotista             string
	ValorTotal             string
	CreatedAt              time.Time
	UpdateAt               time.Time
}
