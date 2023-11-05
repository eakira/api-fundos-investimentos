package domain

import "time"

type ArquivosDomain struct {
	Id          string
	Endereco    string
	TipoArquivo string
	Referencia  string
	Status      string
	Download    bool
	Processado  bool
	CreatedAt   time.Time
	UpdateAt    time.Time
}
