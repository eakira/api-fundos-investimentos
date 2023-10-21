package dto

type FundosQueueDto struct {
	Topic string `json:"Topic"`
	Queue string `json:"queue"`
	Data  []any  `json:"data"`
}
