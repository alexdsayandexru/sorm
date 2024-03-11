package validator

import (
	"errors"
	"fmt"
)

func (v *Validator) Maximum(value int32) *Validator {
	if v.error != nil {
		return v
	}
	target, ok := v.target.(int32)
	if ok == false {
		return v
	}
	if target > value {
		v.error = errors.New(fmt.Sprintf(MaxValue, value))
	}
	return v
}

func (v *Validator) Equal(value int32) *Validator {
	if v.error != nil {
		return v
	}
	target, ok := v.target.(int32)
	if ok == false {
		return v
	}
	if target != value {
		v.error = errors.New(fmt.Sprintf(EqualValue, value))
	}
	return v
}
