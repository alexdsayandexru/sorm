package models

type UserAccountRecoveryEventData struct {
	Datetime  string `json:"datetime"`
	ServiceId string `json:"service_id"`
	UserId    string `json:"user_id"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	Msisdn    string `json:"msisdn"`
	Email     string `json:"email"`
	UserAgent string `json:"user_agent"`
}

type UserAccountRecovery struct {
	EventType     string                       `json:"event_type"`
	CorrelationId string                       `json:"correlation_id"`
	TelcoId       int                          `json:"telco_id"`
	UserType      int                          `json:"user_type"`
	EventData     UserAccountRecoveryEventData `json:"event_data"`
}
