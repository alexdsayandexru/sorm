package models

type LoginUserEvent struct {
	ServiceId int    `json:"service_id"`
	UserId    string `json:"user_id"`
	Msisdn    string `json:"msisdn"`
	Email     string `json:"email"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	UserAgent string `json:"user_agent"`
	Factor    string `json:"factor"`
	Datetime  string `json:"datetime"`
}

type LoginUser struct {
	EventType     string         `json:"event_type"`
	CorrelationId string         `json:"correlation_id"`
	TelcoId       int            `json:"telco_id"`
	UserType      int            `json:"user_type"`
	Event         LoginUserEvent `json:"event"`
}
