package requestValidator

import (
	"github.com/go-playground/validator/v10"
)

// Comment
type CustomValidator struct {
	Validator *validator.Validate
}

// implement echo validator
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		return err
	}

	return nil
}
