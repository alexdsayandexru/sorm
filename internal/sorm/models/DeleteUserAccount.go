package models

type DeleteUserAccountEventData struct {
	ServiceId int32  `json:"service_id"`
	UserId    string `json:"user_id"`
	Ip        string `json:"ip"`
	Port      int32  `json:"port"`
	UserAgent string `json:"user_agent"`
	Message   string `json:"message"`
	Datetime  string `json:"datetime"`
}

type DeleteUserAccount struct {
	EventType     string                     `json:"event_type"`
	CorrelationId string                     `json:"correlation_id"`
	TelcoId       int32                      `json:"telco_id"`
	UserType      int32                      `json:"user_type"`
	EventData     DeleteUserAccountEventData `json:"event_data"`
}

func (target *DeleteUserAccount) GetRules() ValidationRules {
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
