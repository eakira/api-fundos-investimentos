package request

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Date time.Time

type Cnpj string

type Decimal float64

// Implementação da interface json.Unmarshaler
func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	// Parse da string da data para time.Time
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic("teste")
		return err
	}

	*d = Date(parsedDate)
	return nil
}

// Implementação da interface json.Unmarshaler
func (c *Cnpj) UnmarshalJSON(data []byte) error {
	rawCnpj := strings.Trim(string(data), `"`)

	mapFunc := func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1 // Retorna -1 para remover o caractere
	}
	// Atribuir o valor parseado ao campo Date
	*c = Cnpj(strings.Map(mapFunc, rawCnpj))

	return nil
}

// Implementação da interface json.Unmarshaler
func (d *Decimal) UnmarshalJSON(data []byte) error {
	rawNumber := strings.Trim(string(data), `"`)

	// Converter a string para um float64
	number, err := strconv.ParseFloat(rawNumber, 64)
	if err != nil {
		return fmt.Errorf("falha ao converter número: %v", err)
	}

	// Atribuir o valor convertido ao campo Number
	*d = Decimal(number)

	return nil
}
