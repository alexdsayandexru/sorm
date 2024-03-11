package models

import "github.com/alexdsayandexru/sorm/internal/validator"

type UpdateUserDataAbInfo struct {
	UserId          string       `json:"user_id"`
	Nick            string       `json:"nick"`
	BirthDate       string       `json:"birth_date"`
	Name            string       `json:"name"`
	Family          string       `json:"family"`
	Initial         string       `json:"initial"`
	Msisdns         []string     `json:"msisdn"`
	Emails          []string     `json:"email"`
	Address         string       `json:"address"`
	DatetimeReg     string       `json:"datetime_reg"`
	DatetimeUpdated string       `json:"datetime_updated"`
	ServiceId       int32        `json:"service_id"`
	ContractDate    string       `json:"contract_date"`
	ImId            []ImId       `json:"im_id"`
	Additional      []Additional `json:"additional"`
}

type UpdateUserDataEventData struct {
	ServiceId int32  `json:"service_id"`
	UserId    string `json:"user_id"`
	Ip        string `json:"ip"`
	Port      int32  `json:"port"`
	UserAgent string `json:"user_agent"`
	Datetime  string `json:"datetime"`
	Message   string `json:"message"`
}

type UpdateUserData struct {
	EventType     string                  `json:"event_type"`
	CorrelationId string                  `json:"correlation_id"`
	TelcoId       int32                   `json:"telco_id"`
	UserType      int32                   `json:"user_type"`
	AbInfo        UpdateUserDataAbInfo    `json:"ab_info"`
	EventData     UpdateUserDataEventData `json:"event_data"`
}

func (target *UpdateUserData) GetRules() ValidationRules {
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
				for i, m := range target.AbInfo.Msisdns {
					ok, err := validator.Validate(m).MaxLength(16).Regex("^[0-9+]+$").GetResult()
					if !ok {
						return false, NewError(i, "msisdn", err)
					}
				}
			}
			return ok, err
		},
		"emails": func() (bool, error) {
			ok, err := validator.Validate(target.AbInfo.Emails).RequiredIf(target.AbInfo.Msisdns == nil || len(target.AbInfo.Msisdns) == 0).GetResult()
			if ok {
				for i, m := range target.AbInfo.Emails {
					ok, err := validator.Validate(m).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
					if !ok {
						return false, NewError(i, "email", err)
					}
				}
			}
			return ok, err
		},
		"nick": func() (bool, error) {
			return validator.Validate(target.AbInfo.Nick).MaxLength(150).GetResult()
		},
		"birth_date": func() (bool, error) {
			return validator.Validate(target.AbInfo.BirthDate).Length(10).Regex("^[0-9-]+$").GetResult()
		},
		"name": func() (bool, error) {
			return validator.Validate(target.AbInfo.Name).MaxLength(30).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
		},
		"family": func() (bool, error) {
			return validator.Validate(target.AbInfo.Family).MaxLength(30).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
		},
		"initial": func() (bool, error) {
			return validator.Validate(target.AbInfo.Initial).MaxLength(30).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
		},
		"address": func() (bool, error) {
			return validator.Validate(target.AbInfo.Address).MaxLength(255).GetResult()
		},
		"im_id": func() (bool, error) {
			if target.AbInfo.ImId == nil {
				return true, nil
			}

			for i, m := range target.AbInfo.ImId {
				ok, err := validator.Validate(m.ServiceId).MaxLength(255).GetResult()
				if !ok {
					return false, NewError(i, "service_id", err)
				}
				ok, err = validator.Validate(m.ServiceName).MaxLength(255).GetResult()
				if !ok {
					return false, NewError(i, "service_name", err)
				}
			}
			return true, nil
		},
		"datetime_reg": func() (bool, error) {
			return validator.Validate(target.AbInfo.DatetimeReg).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		"datetime_updated": func() (bool, error) {
			return validator.Validate(target.AbInfo.DatetimeUpdated).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		"service_user": func() (bool, error) {
			return validator.Validate(target.AbInfo.ServiceId).Required().Equal(1).GetResult()
		},
		"additional": func() (bool, error) {
			if target.AbInfo.Additional == nil {
				return true, nil
			}

			for i, m := range target.AbInfo.Additional {
				ok, err := validator.Validate(m.Title).MaxLength(255).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
				if !ok {
					return false, NewError(i, "title", err)
				}
				ok, err = validator.Validate(m.Content).MaxLength(255).GetResult()
				if !ok {
					return false, NewError(i, "content", err)
				}
			}
			return true, nil
		},
		"contract_date": func() (bool, error) {
			return validator.Validate(target.AbInfo.ContractDate).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		"service_id": func() (bool, error) {
			return validator.Validate(target.EventData.ServiceId).Required().Maximum(100000000).GetResult()
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
