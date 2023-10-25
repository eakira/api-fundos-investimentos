package domain

import "time"

type ArquivosDomain struct {
	Endereco    string
	TipoArquivo string
	Referencia  string
	Status      string
	Download    bool
	Processado  bool
	CreatedAt   time.Time
	UpdateAt    time.Time
}
