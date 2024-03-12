package models

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
		"datetime_reg": func() (bool, error) {
			return ValidateRequiredDatetime(target.AbInfo.DatetimeReg)
		},
		"service_user": func() (bool, error) {
			return ValidateServiceUser(target.AbInfo.ServiceId)
		},
		"contract_date": func() (bool, error) {
			return ValidateRequiredDatetime(target.AbInfo.ContractDate)
		},
		"service_id": func() (bool, error) {
			return ValidateServiceId(target.EventData.ServiceId)
		},
		"msisdn_login": func() (bool, error) {
			return ValidateMsisdnLogin(target.EventData.Msisdn, target.EventData.Email)
		},
		"email_login": func() (bool, error) {
			return ValidateEmailLogin(target.EventData.Email, target.EventData.Msisdn)
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
