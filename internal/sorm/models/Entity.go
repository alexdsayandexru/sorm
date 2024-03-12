package models

import "encoding/json"

type IEntity interface {
	GetRules() ValidationRules
}

func ToBytes(entity IEntity) []byte {
	bites, _ := json.Marshal(entity)
	return bites
}
