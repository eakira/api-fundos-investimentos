package request

type FundosQueueSincronizarRequest struct {
	Queue string   `json:"queue"`
	Data  []string `json:"data"`
}

type FundosDownloadCvmFilesQueueRequest struct {
	FileName   string `json:"file-name"`
	Referencia string `json:"referencia"`
	Tipo       string `json:"tipo"`
}
