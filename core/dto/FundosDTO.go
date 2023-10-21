package dto

type FundosQueueDto struct {
	Topic string `json:"-"`
	Queue string `json:"queue"`
	Data  []any  `json:"data"`
}
