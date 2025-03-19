package utils

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

func NewErrorResponse(message string, field string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Field:   field,
	}
}

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var element ErrorResponse

			element.Field = err.Field()
			element.Message = err.Error()

			validationErrors = append(validationErrors, element)
		}

	}

	return validationErrors
}
