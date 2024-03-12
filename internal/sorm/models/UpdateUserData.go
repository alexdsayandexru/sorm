package models

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
		"service_user": func() (bool, error) {
			return ValidateServiceUser(target.AbInfo.ServiceId)
		},
		"additional": func() (bool, error) {
			return ValidateAdditional(target.AbInfo.Additional)
		},
		"contract_date": func() (bool, error) {
			return ValidateRequiredDatetime(target.AbInfo.ContractDate)
		},
		"service_id": func() (bool, error) {
			return ValidateServiceId(target.EventData.ServiceId)
		},
		"ip": func() (bool, error) {
			return ValidateIp(target.EventData.Ip)
		},
		"port": func() (bool, error) {
			return ValidatePort(target.EventData.Port)
		},
		"user_agent": func() (bool, error) {
			return ValidateUserAgent(target.EventData.UserAgent)
		},
		"message": func() (bool, error) {
			return ValidateMessage(target.EventData.Message)
		},
		"datetime": func() (bool, error) {
			return ValidateRequiredDatetime(target.EventData.Datetime)
		},
	}
	return validationRules
}
