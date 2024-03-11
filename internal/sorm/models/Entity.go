package models

type IEntity interface {
	GetRules() ValidationRules
}
