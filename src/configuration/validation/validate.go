package validation

import (
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	pt_BR_translation "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enTranslator := en.New()
		unt := ut.New(enTranslator, enTranslator)
		transl, _ = unt.GetTranslator("en")
		err := pt_BR_translation.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(validation_err error) *resterrors.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterrors.NewBadRequestError("Tipo de campo invalido")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []resterrors.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := resterrors.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return resterrors.NewBadRequestValidationError("Alguns campos são invalído", errorCauses)
	} else {
		return resterrors.NewBadRequestError("Erro ao tentar converter os campos")
	}
}
