package models

import (
	"errors"
	"fmt"
)

type IEntity interface {
	GetRules() ValidationRules
}

func NewError(i int, field string, err error) error {
	return errors.New(fmt.Sprintf("[%d].[%s].[%s]", i, field, err.Error()))
}
