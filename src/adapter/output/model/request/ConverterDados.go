package request

import (
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Date time.Time

type Cnpj string

type Decimal float64

type Integer int

func (d *Date) UnmarshalJSON(data []byte) error {
	dateStr := strings.Trim(string(data), `"`)

	if dateStr == "" {
		return nil
	}

	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		logger.Error("erro ao converter a data", err, "sincronizarFundos")
		return err
	}

	*d = Date(parsedDate)
	return nil
}

func (c *Cnpj) UnmarshalJSON(data []byte) error {
	rawCnpj := strings.Trim(string(data), `"`)

	if rawCnpj == "" {
		return nil
	}

	mapFunc := func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}

	*c = Cnpj(strings.Map(mapFunc, rawCnpj))

	return nil
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	rawNumber := strings.Trim(string(data), `"`)

	if rawNumber == "" {
		return nil
	}

	number, err := strconv.ParseFloat(rawNumber, 64)
	if err != nil {
		logger.Error("falha ao converter número:", err, "sincronizarFundos")
		return err
	}

	*d = Decimal(number)

	return nil
}

func (d *Integer) UnmarshalJSON(data []byte) error {
	rawNumber := strings.Trim(string(data), `"`)

	if rawNumber == "" {
		return nil
	}

	number, err := strconv.ParseInt(rawNumber, 10, 64)
	if err != nil {
		logger.Error("falha ao converter número:", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	*d = Integer(number)

	return nil
}
