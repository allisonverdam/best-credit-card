package errors

import (
	"net/http"
	"sort"

	validation "github.com/go-ozzo/ozzo-validation"
)

type validationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// InternalServerError cria um novo erro representando o "internal server" (HTTP 500)
func InternalServerError(err error) *APIError {
	return NewAPIError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", Params{"error": err.Error()})
}

// NotFound cria um novo erro representando o erro "resource-not-found" (HTTP 404)
func NotFound(resource string) *APIError {
	return NewAPIError(http.StatusNotFound, "NOT_FOUND", Params{"resource": resource})
}

// Unauthorized cria um novo erro representando o erro "authentication failure" (HTTP 401)
func Unauthorized(err string) *APIError {
	return NewAPIError(http.StatusUnauthorized, "UNAUTHORIZED", Params{"error": err})
}

// InvalidData converte o erro "data validation" em um erro customizado (HTTP 400)
func InvalidData(errs validation.Errors) *APIError {
	result := []validationError{}
	fields := []string{}
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		err := errs[field]
		result = append(result, validationError{
			Field: field,
			Error: err.Error(),
		})
	}

	err := NewAPIError(http.StatusBadRequest, "INVALID_DATA", nil)
	err.Details = result

	return err
}