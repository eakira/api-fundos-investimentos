package response

type FundosQueueResponse struct {
	Topic string   `json:"-"`
	Queue string   `json:"queue"`
	Data  []string `json:"data"`
}
