package validation

import (
	"encoding/json"
	"errors"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *rest_err.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		var errorsCauses []rest_err.Cause

		for _, err := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Cause{
				Message: err.Translate(transl),
				Field:   err.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
