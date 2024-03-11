package validator

import (
	"github.com/google/uuid"
)

func (v *Validator) Uiid() *Validator {
	if v.error != nil {
		return v
	}

	value, ok := v.target.(string)
	if !ok {
		panic("Unexpected target type")
	}

	_, err := uuid.Parse(value)
	if err != nil {
		v.error = err
	}

	return v
}
