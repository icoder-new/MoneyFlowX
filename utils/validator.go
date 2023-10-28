package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"strings"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationError(err error) []ErrorMessage {
	errorMessages := []ErrorMessage{}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		e := strings.Split(err.Error(), ".")[1]
		errors := strings.Split(e, " of type ")
		message := fmt.Sprintf("%s must be %s", errors[0], errors[1])
		errorMessages = append(errorMessages, ErrorMessage{
			Field:   errors[0],
			Message: message,
		})
		return errorMessages
	}

	if errors.Is(err, io.EOF) {
		for _, e := range err.(validator.ValidationErrors) {
			var message string

			if e.Param() != "" {
				message = ToSnakeCase(e.Field()) + " " + e.Tag() + " " + e.Param()
			} else {
				message = ToSnakeCase(e.Field()) + " " + e.Tag()
			}
			errorMessages = append(errorMessages, ErrorMessage{
				Field:   ToSnakeCase(e.Field()),
				Message: message,
			})
		}
		return errorMessages
	}

	return errorMessages
}
