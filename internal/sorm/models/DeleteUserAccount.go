package models

type DeleteUserAccountEventData struct {
	Datetime  string `json:"datetime"`
	ServiceId string `json:"service_id"`
	UserId    string `json:"user_id"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	UserAgent string `json:"user_agent"`
	Message   string `json:"message"`
}

type DeleteUserAccount struct {
	EventType     string                     `json:"event_type"`
	CorrelationId string                     `json:"correlation_id"`
	TelcoId       int                        `json:"telco_id"`
	UserType      int                        `json:"user_type"`
	EventData     DeleteUserAccountEventData `json:"event_data"`
}
