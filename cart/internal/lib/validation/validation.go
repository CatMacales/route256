package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"sync"
)

var (
	globalValidator *Validator
	once            = sync.Once{}
)

type Validator struct {
	v *validator.Validate
}

// InitValidator initializes and returns a singleton instance of the Validator.
// The Validator is configured with required struct validation enabled.
func InitValidator() *Validator {
	once.Do(func() {
		globalValidator = &Validator{v: validator.New(validator.WithRequiredStructEnabled())}
	})
	return globalValidator
}

func BeautyStructValidate(s any) error {
	if globalValidator == nil {
		panic("global validator is nil")
	}

	rawErr := globalValidator.v.Struct(s)
	if rawErr != nil {
		var err validator.ValidationErrors
		if errors.As(rawErr, &err) {
			var errorMessages []string
			for _, validationErr := range err {
				switch validationErr.Tag() {
				case "gte":
					errorMessages = append(errorMessages, fmt.Sprintf("%s must be greater than %s", validationErr.Field(), validationErr.Param()))
				case "required":
					errorMessages = append(errorMessages, fmt.Sprintf("%s is required", validationErr.Field()))
				default:
					errorMessages = append(errorMessages, validationErr.Error())
				}
			}
			return fmt.Errorf(strings.Join(errorMessages, ", "))
		}
		return rawErr
	}
	return nil
}
