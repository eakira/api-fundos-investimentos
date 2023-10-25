package response

type FundosQueueResponse struct {
	Topic string   `json:"-"`
	Queue string   `json:"queue"`
	Data  []string `json:"data"`
}

type FundosDownloadCvmFilesQueueResponse struct {
	FileName   string `json:"file-name"`
	Referencia string `json:"referencia"`
	Tipo       string `json:"tipo"`
}
