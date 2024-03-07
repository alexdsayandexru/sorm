package models

import (
	"github.com/alexdsayandexru/sorm/internal/validator"
)

type AbInfo struct {
	UserId       string       `json:"user_id"`
	Msisdns      []string     `json:"msisdn"`
	Emails       []string     `json:"email"`
	DatetimeReg  string       `json:"datetime_reg"`
	ServiceId    int32        `json:"service_id"`
	Additional   []Additional `json:"additional"`
	ContractDate string       `json:"contract_date"`
}

type RegisterUserEventData struct {
	ServiceId int32  `json:"service_id"`
	UserId    string `json:"user_id"`
	Msisdn    string `json:"msisdn"`
	Email     string `json:"email"`
	Ip        string `json:"ip"`
	Port      int32  `json:"port"`
	UserAgent string `json:"user_agent"`
	Datetime  string `json:"datetime"`
	Message   string `json:"message"`
}

type RegisterUser struct {
	EventType     string                `json:"event_type"`
	CorrelationId string                `json:"correlation_id"`
	TelcoId       int32                 `json:"telco_id"`
	UserType      int32                 `json:"user_type"`
	AbInfo        AbInfo                `json:"ab_info"`
	EventData     RegisterUserEventData `json:"event_data"`
}

func (target *RegisterUser) GetRules() ValidationRules {
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
			return validator.Validate(target.AbInfo.UserId).Required().MaxLength(255).Regex("^[A-Za-z0-9_-]+$").GetResult()
		},
		"msisdns": func() (bool, error) {
			ok, err := validator.Validate(target.AbInfo.Msisdns).RequiredIf(target.AbInfo.Emails == nil || len(target.AbInfo.Emails) == 0).GetResult()
			if ok {
				for _, m := range target.AbInfo.Msisdns {
					ok, err := validator.Validate(m).MaxLength(16).Regex("^[0-9+]+$").GetResult()
					if !ok {
						return ok, err
					}
				}
			}
			return ok, err
		},
		"emails": func() (bool, error) {
			ok, err := validator.Validate(target.AbInfo.Emails).RequiredIf(target.AbInfo.Msisdns == nil || len(target.AbInfo.Msisdns) == 0).GetResult()
			if ok {
				for _, m := range target.AbInfo.Emails {
					ok, err := validator.Validate(m).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
					if !ok {
						return ok, err
					}
				}
			}
			return ok, err
		},
		"datetime_reg": func() (bool, error) {
			return validator.Validate(target.AbInfo.DatetimeReg).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		/*"service_user": func() (bool, error) {
			return validator.Validate(target.ServiceUser).Required().Equal(1).GetResult()
		},*/
		"contract_date": func() (bool, error) {
			return validator.Validate(target.AbInfo.ContractDate).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		"service_id": func() (bool, error) {
			return validator.Validate(target.AbInfo.ServiceId).Required().Maximum(100000000).GetResult()
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
			return validator.Validate(target.EventData.UserAgent).MaxLength(1023).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
		},
		"message": func() (bool, error) {
			return validator.Validate(target.EventData.Message).Required().MaxLength(8000).GetResult()
		},
		"datetime": func() (bool, error) {
			return validator.Validate(target.EventData.Datetime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
	}
	return validationRules
}
