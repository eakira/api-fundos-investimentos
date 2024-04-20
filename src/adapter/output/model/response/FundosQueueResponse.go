package response

import "time"

type FundosQueueResponse struct {
	Topic string
	Queue string
	Data  []byte
}

type FundosDownloadCvmFilesQueueResponse struct {
	Id          string    `json:"id"`
	Endereco    string    `json:"endereco"`
	TipoArquivo string    `json:"tipo-arquivo"`
	Referencia  string    `json:"referencia"`
	Status      string    `json:"status"`
	Download    string    `json:"download"`
	Processado  string    `json:"processado"`
	Baixar      bool      `json:"baixar"`
	CreatedAt   time.Time `json:"created-at"`
	UpdatedAt   time.Time `json:"update-at"`
}
