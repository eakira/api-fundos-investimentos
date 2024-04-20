package domain

import "time"

type InformacaoDiariaDomain struct {
	Id                     string
	FundoCnpj              string
	CaptacaoDia            float64
	DataCompetencia        time.Time
	ResgateDia             float64
	TipoFundo              string
	ValorPatrimonioLiquido float64
	CotaValor              float64
	QtdCotista             int
	ValorTotal             float64
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
