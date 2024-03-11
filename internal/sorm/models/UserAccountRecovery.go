package models

import "github.com/alexdsayandexru/sorm/internal/validator"

type UserAccountRecoveryEventData struct {
	Datetime  string `json:"datetime"`
	ServiceId int32  `json:"service_id"`
	UserId    string `json:"user_id"`
	Ip        string `json:"ip"`
	Port      int32  `json:"port"`
	Msisdn    string `json:"msisdn"`
	Email     string `json:"email"`
	UserAgent string `json:"user_agent"`
}

type UserAccountRecovery struct {
	EventType     string                       `json:"event_type"`
	CorrelationId string                       `json:"correlation_id"`
	TelcoId       int32                        `json:"telco_id"`
	UserType      int32                        `json:"user_type"`
	EventData     UserAccountRecoveryEventData `json:"event_data"`
}

func (target *UserAccountRecovery) GetRules() ValidationRules {
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
			return validator.Validate(target.EventData.UserId).Required().MaxLength(255).Regex("^[A-Za-z0-9_-]+$").GetResult()
		},
		"service_id": func() (bool, error) {
			return validator.Validate(target.EventData.ServiceId).Required().Maximum(100000000).GetResult()
		},
		"msisdn_login": func() (bool, error) {
			return validator.Validate(target.EventData.Msisdn).RequiredIf(len(target.EventData.Email) == 0).MaxLength(16).Regex("^[0-9+]+$").GetResult()
		},
		"email_login": func() (bool, error) {
			return validator.Validate(target.EventData.Email).RequiredIf(len(target.EventData.Msisdn) == 0).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
		},
		"ip": func() (bool, error) {
			return validator.Validate(target.EventData.Ip).MaxLength(255).Regex("^[0-9.]+$").GetResult()
		},
		"port": func() (bool, error) {
			return validator.Validate(target.EventData.Port).Maximum(99999).GetResult()
		},
		"user_agent": func() (bool, error) {
			return validator.Validate(target.EventData.UserAgent).MaxLength(1023).Regex("^[A-Za-z -]+$").GetResult()
		},
		"datetime": func() (bool, error) {
			return validator.Validate(target.EventData.Datetime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
	}
	return validationRules
}
