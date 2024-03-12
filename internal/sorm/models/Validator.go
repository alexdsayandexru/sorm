package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ValidationRules map[string]func() (bool, error)

const (
	ErrorsCode        = "field-errors"
	ErrorsDescription = "Ошибка валидации данных"
)

type FieldError struct {
	FieldName        string `json:"field_name"`
	ErrorDescription string `json:"error_description"`
}

type ErrorInfo struct {
	ErrorsCode        string       `json:"errors_code"`
	ErrorsDescription string       `json:"errors_description"`
	FieldErrors       []FieldError `json:"field_errors"`
}

func (e *ErrorInfo) ToBytes() []byte {
	bites, _ := json.Marshal(*e)
	return bites
}

func (e *ErrorInfo) ToJson() string {
	return string(e.ToBytes())
}

type IValidator interface {
	Validate(entity IEntity) (bool, ErrorInfo)
}

func GetValidator() *ValidatorImpl {
	return &ValidatorImpl{}
}

type ValidatorImpl struct {
}

func (v *ValidatorImpl) Validate(entity IEntity) (bool, ErrorInfo) {
	errorInfo := ErrorInfo{
		ErrorsCode:        ErrorsCode,
		ErrorsDescription: ErrorsDescription,
		FieldErrors:       []FieldError{},
	}
	for key, rule := range entity.GetRules() {
		ok, err := rule()
		if !ok {
			errorInfo.FieldErrors = append(errorInfo.FieldErrors, FieldError{FieldName: key, ErrorDescription: err.Error()})
		}
	}
	return len(errorInfo.FieldErrors) == 0, errorInfo
}

func NewError(i int, field string, err error) error {
	return errors.New(fmt.Sprintf("[%d].[%s].[%s]", i, field, err.Error()))
}
