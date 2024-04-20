package request

import "time"

type FundosQueueSincronizarRequest struct {
	Queue string   `json:"queue"`
	Data  []string `json:"data"`
}

type FundosCvmArquivosQueueRequest struct {
	Id          string    `json:"id"`
	Endereco    string    `json:"endereco"`
	TipoArquivo string    `json:"tipo-arquivo"`
	Referencia  string    `json:"referencia"`
	Status      string    `json:"status"`
	Download    string    `json:"download"`
	Processado  string    `json:"processado"`
	CreatedAt   time.Time `json:"created-at"`
	UpdatedAt   time.Time `json:"update-at"`
}
