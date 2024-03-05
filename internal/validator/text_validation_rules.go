package validator

import (
	"errors"
	"fmt"
	"regexp"
)

func (v *Validator) MaxLength(max int) *Validator {
	if v.error != nil {
		return v
	}
	value, ok := v.target.(string)
	if ok == false {
		return v
	}
	if len(value) > max {
		v.error = errors.New(fmt.Sprintf(MaxFieldLength, max))
	}
	return v
}

func (v *Validator) Length(count int) *Validator {
	if v.error != nil {
		return v
	}
	value, ok := v.target.(string)
	if ok == false {
		return v
	}
	if len(value) != count {
		v.error = errors.New(fmt.Sprintf(EqualFieldLength, count))
	}
	return v
}

func (v *Validator) Regex(expression string) *Validator {
	if v.error != nil {
		return v
	}
	value, ok := v.target.(string)
	if ok == false {
		return v
	}
	re := regexp.MustCompile(expression)
	if !re.MatchString(value) {
		v.error = errors.New(fmt.Sprintf(InvalidValue))
	}
	return v
}
