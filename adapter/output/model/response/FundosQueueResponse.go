package response

import "time"

type FundosQueueResponse struct {
	Topic string
	Queue string
	Data  []byte
}

type FundosDownloadCvmFilesQueueResponse struct {
	Endereco    string    `json:"endereco"`
	TipoArquivo string    `json:"tipo-arquivo"`
	Referencia  string    `json:"referencia"`
	Status      string    `json:"status"`
	Download    string    `json:"download"`
	Processado  string    `json:"processado"`
	CreatedAt   time.Time `json:"created-at"`
	UpdateAt    time.Time `json:"update-at"`
}
