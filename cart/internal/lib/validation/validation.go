package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"sync"
)

var (
	globalValidator *Validator
	once            = sync.Once{}
)

type Validator struct {
	v *validator.Validate
}

func NewValidator() {
	once.Do(func() {
		globalValidator = &Validator{v: validator.New(validator.WithRequiredStructEnabled())}
	})
}

func BeautyStructValidate(s any) error {
	if globalValidator == nil {
		panic("global validator is nil")
	}

	rawErr := globalValidator.v.Struct(s)
	if rawErr != nil {
		var err validator.ValidationErrors
		if errors.As(rawErr, &err) {
			switch err[0].Tag() {
			case "gte":
				return fmt.Errorf("%s must be greater than %s", err[0].Field(), err[0].Param())
			case "required":
				return fmt.Errorf("%s is required", err[0].Field())
			default:
				return err
			}
		}
		return rawErr
	}
	return nil
}
