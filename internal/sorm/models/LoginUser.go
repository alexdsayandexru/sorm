package models

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
			return ValidateCorrelationId(target.CorrelationId)
		},
		"telco_id": func() (bool, error) {
			return ValidateTelcoId(target.TelcoId)
		},
		"user_type": func() (bool, error) {
			return ValidateUserType(target.UserType)
		},
		"user_id": func() (bool, error) {
			return ValidateUserId(target.Event.UserId)
		},
		"service_id": func() (bool, error) {
			return ValidateServiceId(target.Event.ServiceId)
		},
		"msisdn_login": func() (bool, error) {
			return ValidateMsisdnLogin(target.Event.Msisdn, target.Event.Email)
		},
		"email_login": func() (bool, error) {
			return ValidateEmailLogin(target.Event.Email, target.Event.Msisdn)
		},
		"ip": func() (bool, error) {
			return ValidateIp(target.Event.Ip)
		},
		"port": func() (bool, error) {
			return ValidatePort(target.Event.Port)
		},
		"user_agent": func() (bool, error) {
			return ValidateUserAgent(target.Event.UserAgent)
		},
		"factor": func() (bool, error) {
			return ValidateFactor(target.Event.Factor)
		},
		"datetime": func() (bool, error) {
			return ValidateRequiredDatetime(target.Event.Datetime)
		},
	}
	return validationRules
}
