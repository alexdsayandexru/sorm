package validator

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func (v *Validator) Uiid() *Validator {
	if v.error != nil {
		return v
	}
	if v.target == nil {
		return v
	}
	value, ok := v.target.(string)
	if ok == false {
		v.error = errors.New(fmt.Sprintf(InvalidValue))
	} else {
		_, err := uuid.Parse(value)
		if err != nil {
			v.error = err
		}
	}
	return v
}
