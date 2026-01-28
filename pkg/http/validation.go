package httpsuite

import (
	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

// IsRequestValid validates the provided request struct using the go-playground/validator package.
// It returns a ProblemDetails instance if validation fails, or nil if the request is valid.
func IsRequestValid(request any) error {
	err := validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}