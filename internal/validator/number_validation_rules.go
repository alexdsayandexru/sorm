package validator

import (
	"errors"
	"fmt"
)

func (v *Validator) Maximum(max int) *Validator {
	if v.error != nil {
		return v
	}
	target, ok := v.target.(int)
	if ok == false {
		return v
	}
	if target > max {
		v.error = errors.New(fmt.Sprintf(MaxValue, max))
	}
	return v
}

func (v *Validator) Equal(value int) *Validator {
	if v.error != nil {
		return v
	}
	target, ok := v.target.(int)
	if ok == false {
		return v
	}
	if target != value {
		v.error = errors.New(fmt.Sprintf(EqualValue, value))
	}
	return v
}
