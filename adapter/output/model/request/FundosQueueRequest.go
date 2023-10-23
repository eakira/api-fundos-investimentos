package request

type FundosQueueSincronizarRequest struct {
	Queue string   `json:"queue"`
	Data  []string `json:"data"`
}
