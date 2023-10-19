package validation

import (
	"api-fundos-investimentos/configuration/rest_errors"
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

func ValidateUserError(validation_err error) *rest_errors.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_errors.NewBadRequestError("Tipo de campo invalido")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_errors.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_errors.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_errors.NewBadRequestValidationError("Alguns campos são invalído", errorCauses)
	} else {
		return rest_errors.NewBadRequestError("Erro ao tentar converter os campos")
	}
}
