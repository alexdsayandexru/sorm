package models

import (
	"github.com/alexdsayandexru/sorm/internal/validator"
)

type Service struct {
	ServiceId   int32  `json:"service_id"`
	Description string `json:"description"`
	BeginTime   string `json:"begin_time"`
	EndTime     string `json:"end_time"`
}

type DirectoryData struct {
	Datetime      string    `json:"datetime"`
	EventType     string    `json:"event_type"`
	CorrelationId string    `json:"correlation_id"`
	OriServices   []Service `json:"oriServices"`
}

func (target *DirectoryData) GetRules() ValidationRules {
	validationRules := map[string]func() (bool, error){
		"correlation_id": func() (bool, error) {
			return validator.Validate(target.CorrelationId).Required().Uiid().GetResult()
		},
		"datetime": func() (bool, error) {
			return validator.Validate(target.Datetime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		"oriServices": func() (bool, error) {
			for i, s := range target.OriServices {
				ok, err := validator.Validate(s.ServiceId).Required().Maximum(100000000).GetResult()
				if !ok {
					return false, NewError(i, "service_id", err)
				}
				ok, err = validator.Validate(s.Description).MaxLength(255).GetResult()
				if !ok {
					return false, NewError(i, "description", err)
				}
				ok, err = validator.Validate(s.BeginTime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
				if !ok {
					return false, NewError(i, "begin_time", err)
				}
				ok, err = validator.Validate(s.EndTime).Length(23).Regex("^[0-9 :.-]+$").GetResult()
				if !ok {
					return false, NewError(i, "end_time", err)
				}
			}
			return true, nil
		},
	}
	return validationRules
}
