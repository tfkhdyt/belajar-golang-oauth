package validator

import (
	"github.com/go-playground/validator/v10"

	"github.com/tfkhdyt/belajar-golang-oauth/pkg/errs"
)

var validate = validator.New()

type Validator struct{}

func (v *Validator) Validate() []errs.ErrorResponse {
	var errors []errs.ErrorResponse
	if err := validate.Struct(v); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errs.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}

	return errors
}
