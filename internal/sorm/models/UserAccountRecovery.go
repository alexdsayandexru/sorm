package models

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
			return ValidateCorrelationId(target.CorrelationId)
		},
		"telco_id": func() (bool, error) {
			return ValidateTelcoId(target.TelcoId)
		},
		"user_type": func() (bool, error) {
			return ValidateUserType(target.UserType)
		},
		"user_id": func() (bool, error) {
			return ValidateUserId(target.EventData.UserId)
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
		"datetime": func() (bool, error) {
			return ValidateRequiredDatetime(target.EventData.Datetime)
		},
	}
	return validationRules
}
