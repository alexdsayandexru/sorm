package models

type UpdateUserDataAbInfoAdditional struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateUserDataAbInfoImId struct {
	ServiceName string `json:"service_name"`
	ServiceId   string `json:"service_id"`
}

type UpdateUserDataAbInfo struct {
	UserId          string                           `json:"user_id"`
	Nick            string                           `json:"nick"`
	BirthDate       string                           `json:"birth_date"`
	Name            string                           `json:"name"`
	Family          string                           `json:"family"`
	Initial         string                           `json:"initial"`
	Msisdn          []string                         `json:"msisdn"`
	Email           []string                         `json:"email"`
	Address         string                           `json:"address"`
	DatetimeReg     string                           `json:"datetime_reg"`
	DatetimeUpdated string                           `json:"datetime_updated"`
	ServiceId       int                              `json:"service_id"`
	ContractDate    string                           `json:"contract_date"`
	ImId            []UpdateUserDataAbInfoImId       `json:"im_id"`
	Additional      []UpdateUserDataAbInfoAdditional `json:"additional"`
}

type UpdateUserDataEventData struct {
	ServiceId string `json:"service_id"`
	UserId    string `json:"user_id"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	UserAgent string `json:"user_agent"`
	Datetime  string `json:"datetime"`
	Message   string `json:"message"`
}

type UpdateUserData struct {
	EventType     string                  `json:"event_type"`
	CorrelationId string                  `json:"correlation_id"`
	TelcoId       int                     `json:"telco_id"`
	UserType      int                     `json:"user_type"`
	AbInfo        UpdateUserDataAbInfo    `json:"ab_info"`
	EventData     UpdateUserDataEventData `json:"event_data"`
}
