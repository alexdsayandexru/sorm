package models

import "github.com/alexdsayandexru/sorm/internal/validator"

type DeleteAccountAbInfo struct {
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
	ImId            []ImId       `json:"im_id"`
	Additional      []Additional `json:"additional"`
	DatetimeUnreg   string       `json:"datetime_unreg"`
	ContractDate    string       `json:"contract_date"`
}

type DeleteAccount struct {
	EventType     string              `json:"event_type"`
	CorrelationId string              `json:"correlation_id"`
	TelcoId       int32               `json:"telco_id"`
	UserType      int32               `json:"user_type"`
	AbInfo        DeleteAccountAbInfo `json:"ab_info"`
}

func (target *DeleteAccount) GetRules() ValidationRules {
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
		"msisdn": func() (bool, error) {
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
		"email": func() (bool, error) {
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
		"service_id-service_name": func() (bool, error) {
			if target.AbInfo.ImId == nil || len(target.AbInfo.ImId) == 0 {
				return true, nil
			}
			for _, m := range target.AbInfo.ImId {
				ok, err := validator.Validate(m.ServiceId).MaxLength(255).GetResult()
				if !ok {
					return ok, err
				}
				ok, err = validator.Validate(m.ServiceName).MaxLength(255).GetResult()
				if !ok {
					return ok, err
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
		"datetime_unreg": func() (bool, error) {
			return validator.Validate(target.AbInfo.DatetimeUnreg).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		/*"service_user": func() (bool, error) {
			return validator.Validate(target.ServiceUser).Required().Equal(1).GetResult()
		},*/
		"title-content": func() (bool, error) {
			for _, m := range target.AbInfo.Additional {
				ok, err := validator.Validate(m.Title).MaxLength(255).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
				if !ok {
					return false, err
				}
				ok, err = validator.Validate(m.Content).MaxLength(255).GetResult()
				if !ok {
					return false, err
				}
			}
			return true, nil
		},
		"contract_date": func() (bool, error) {
			return validator.Validate(target.AbInfo.ContractDate).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		"service_id": func() (bool, error) {
			return validator.Validate(target.AbInfo.ServiceId).Required().Maximum(100000000).GetResult()
		},
	}
	return validationRules
}
