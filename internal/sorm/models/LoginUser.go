package models

import "github.com/alexdsayandexru/sorm/internal/validator"

type LoginUserEvent struct {
	ServiceId int32  `json:"service_id"`
	UserId    string `json:"user_id"`
	Msisdn    string `json:"msisdn"`
	Email     string `json:"email"`
	Ip        string `json:"ip"`
	Port      int32  `json:"port"`
	UserAgent string `json:"user_agent"`
	Factor    string `json:"factor"`
	Datetime  string `json:"datetime"`
}

type LoginUser struct {
	EventType     string         `json:"event_type"`
	CorrelationId string         `json:"correlation_id"`
	TelcoId       int32          `json:"telco_id"`
	UserType      int32          `json:"user_type"`
	Event         LoginUserEvent `json:"event"`
}

func (target *LoginUser) GetRules() ValidationRules {
	validationRules := map[string]func() (bool, error){
		"correlation_id": func() (bool, error) {
			return validator.Validate(target.CorrelationId).Required().Uiid().GetResult()
		},
		"telco_id": func() (bool, error) {
			return validator.Validate(target.TelcoId).Required().Maximum(100).GetResult()
		},
		"user_type": func() (bool, error) {
			return validator.Validate(target.UserType).Required().Maximum(100).GetResult()
		},
		"user_id": func() (bool, error) {
			return validator.Validate(target.Event.UserId).Required().MaxLength(255).Regex("^[A-Za-z0-9_-]+$").GetResult()
		},
		"service_id": func() (bool, error) {
			return validator.Validate(target.Event.ServiceId).Required().Maximum(100000000).GetResult()
		},
		"msisdn_login": func() (bool, error) {
			return validator.Validate(target.Event.Msisdn).RequiredIf(len(target.Event.Email) == 0).MaxLength(16).Regex("^[0-9+]+$").GetResult()
		},
		"email_login": func() (bool, error) {
			return validator.Validate(target.Event.Email).RequiredIf(len(target.Event.Msisdn) == 0).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
		},
		"ip": func() (bool, error) {
			return validator.Validate(target.Event.Ip).MaxLength(255).Regex("^[0-9.]+$").GetResult()
		},
		"port": func() (bool, error) {
			return validator.Validate(target.Event.Port).Maximum(99999).GetResult()
		},
		"user_agent": func() (bool, error) {
			return validator.Validate(target.Event.UserAgent).MaxLength(1023).Regex("^[A-Za-z -]+$").GetResult()
		},
		"factor": func() (bool, error) {
			return validator.Validate(target.Event.Factor).MaxLength(255).Regex("^[A-Za-z0-9-]+$").GetResult()
		},
		"datetime": func() (bool, error) {
			return validator.Validate(target.Event.Datetime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
	}
	return validationRules
}
