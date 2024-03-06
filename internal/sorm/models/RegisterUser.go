package models

type Additional struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AbInfo struct {
	UserId       string       `json:"user_id"`
	Msisdn       []string     `json:"msisdn"`
	Email        []string     `json:"email"`
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
