package models

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
			return ValidateCorrelationId(target.CorrelationId)
		},
		"telco_id": func() (bool, error) {
			return ValidateTelcoId(target.TelcoId)
		},
		"user_type": func() (bool, error) {
			return ValidateUserType(target.UserType)
		},
		"user_id": func() (bool, error) {
			return ValidateUserId(target.AbInfo.UserId)
		},
		"msisdns": func() (bool, error) {
			return ValidateMsisdns(target.AbInfo.Msisdns, target.AbInfo.Emails)
		},
		"emails": func() (bool, error) {
			return ValidateEmails(target.AbInfo.Emails, target.AbInfo.Msisdns)
		},
		"nick": func() (bool, error) {
			return ValidateNick(target.AbInfo.Nick)
		},
		"birth_date": func() (bool, error) {
			return ValidateBirthDate(target.AbInfo.BirthDate)
		},
		"name": func() (bool, error) {
			return ValidateNameFamilyInitial(target.AbInfo.Name)
		},
		"family": func() (bool, error) {
			return ValidateNameFamilyInitial(target.AbInfo.Family)
		},
		"initial": func() (bool, error) {
			return ValidateNameFamilyInitial(target.AbInfo.Initial)
		},
		"address": func() (bool, error) {
			return ValidateAddress(target.AbInfo.Address)
		},
		"im_id": func() (bool, error) {
			return ValidateImId(target.AbInfo.ImId)
		},
		"datetime_reg": func() (bool, error) {
			return ValidateRequiredDatetime(target.AbInfo.DatetimeReg)
		},
		"datetime_updated": func() (bool, error) {
			return ValidateRequiredDatetime(target.AbInfo.DatetimeUpdated)
		},
		"datetime_unreg": func() (bool, error) {
			return ValidateRequiredDatetime(target.AbInfo.DatetimeUnreg)
		},
		"service_user": func() (bool, error) {
			return ValidateServiceUser(target.AbInfo.ServiceId)
		},
		"additional": func() (bool, error) {
			return ValidateAdditional(target.AbInfo.Additional)
		},
		"contract_date": func() (bool, error) {
			return ValidateRequiredDatetime(target.AbInfo.ContractDate)
		},
	}
	return validationRules
}
