package models

type LogoutUserEvent struct {
	ServiceId int    `json:"service_id"`
	UserId    string `json:"user_id"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	UserAgent string `json:"user_agent"`
	Datetime  string `json:"datetime"`
}

type LogoutUser struct {
	EventType     string          `json:"event_type"`
	CorrelationId string          `json:"correlation_id"`
	TelcoId       int             `json:"telco_id"`
	UserType      int             `json:"user_type"`
	Event         LogoutUserEvent `json:"event"`
}
