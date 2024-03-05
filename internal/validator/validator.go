package validator

import (
	"errors"
)

const (
	InvalidValue     string = "Поле содержит некорректное значение или недопустимые символы"
	RequiredField           = "Поле обязательно для заполнения"
	MaxFieldLength          = "Максимальная длина поля - %d символов"
	EqualFieldLength        = "Длина поля должна быть %d символов"
	MaxValue                = "Поле должно быть меньше или равно %d"
	EqualValue              = "Поле должно быть равно %d"
)

type Validator struct {
	target interface{}
	error  error
}

func Validate(target interface{}) *Validator {
	return &Validator{target: target}
}

func (v *Validator) GetResult() (bool, error) {
	return v.error == nil, v.error
}

func (v *Validator) Required() *Validator {
	if v.error != nil {
		return v
	}
	if v.target == nil {
		v.error = errors.New(RequiredField)
	}
	return v
}

func (v *Validator) RequiredIf(predicate bool) *Validator {
	if v.error != nil {
		return v
	}
	if predicate {
		return v.Required()
	}
	return v
}
