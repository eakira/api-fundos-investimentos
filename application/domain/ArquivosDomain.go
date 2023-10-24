package domain

import "time"

type ArquivosDomain struct {
	endereco    string
	tipoArquivo string
	referencia  string
	status      string
	download    bool
	processado  bool
	createdAt   time.Time
	updateAt    time.Time
}
