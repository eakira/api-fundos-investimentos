package response

type FundosQueueResponse struct {
	Topic string
	Queue string
	Data  []byte
}

type FundosDownloadCvmFilesQueueResponse struct {
	FileName   string `json:"file-name"`
	Referencia string `json:"referencia"`
	Tipo       string `json:"tipo"`
}
