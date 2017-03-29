package validator

import vld "gopkg.in/go-playground/validator.v9"

type requestValidator struct {
	vld *vld.Validate
}

// Validate validate the struct
func (r *requestValidator) Validate(i interface{}) error {
	return r.vld.Struct(i)
}

// NewRequestValidator returns new instance of RequestValidator
func NewRequestValidator() *requestValidator {
	return &requestValidator{vld.New()}
}
