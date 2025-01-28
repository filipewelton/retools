package helpers

import (
	"backend/internal/application/typings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateStruct(data any, err *typings.Error) bool {
	var result = validate.Struct(data)

	if result != nil {
		err.Reason = result.Error()
		return false
	}

	return true
}
