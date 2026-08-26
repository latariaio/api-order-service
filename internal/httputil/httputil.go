package httputil

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// FormatValidationError converte o erro cru do validator em algo legível
func FormatValidationError(err error) []FieldError {
	var validationErrs validator.ValidationErrors
	if !errors.As(err, &validationErrs) {
		return nil
	}

	fieldErrors := make([]FieldError, 0, len(validationErrs))
	for _, fe := range validationErrs {
		fieldErrors = append(fieldErrors, FieldError{
			Field:   toSnakeCase(fe.Field()),
			Message: buildMessage(fe),
		})
	}

	return fieldErrors
}

func buildMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s é obrigatório", fe.Field())
	case "email":
		return fmt.Sprintf("%s deve ser um e-mail válido", fe.Field())
	case "min":
		return fmt.Sprintf("%s deve ter no mínimo %s caracteres", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s deve ter no máximo %s caracteres", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s é inválido", fe.Field())
	}
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
