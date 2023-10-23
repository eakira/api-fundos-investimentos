package request

type FundosQueueRequest struct {
	Topic string   `json:"-"`
	Queue string   `json:"queue"`
	Data  []string `json:"data"`
}
