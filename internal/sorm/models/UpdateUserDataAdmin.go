package models

type UpdateUserDataAdminImId struct {
	ServiceName string `json:"service_name"`
	ServiceId   string `json:"service_id"`
}

type UpdateUserDataAdminAdditional struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateUserDataAdminAbInfo struct {
	UserId          string                          `json:"user_id"`
	Msisdn          []string                        `json:"msisdn"`
	Email           []string                        `json:"email"`
	Nick            string                          `json:"nick"`
	BirthDate       string                          `json:"birth_date"`
	Name            string                          `json:"name"`
	Family          string                          `json:"family"`
	Initial         string                          `json:"initial"`
	Address         string                          `json:"address"`
	DatetimeReg     string                          `json:"datetime_reg"`
	DatetimeUpdated string                          `json:"datetime_updated"`
	ServiceId       int                             `json:"service_id"`
	ImId            []UpdateUserDataAdminImId       `json:"im_id"`
	Additional      []UpdateUserDataAdminAdditional `json:"additional"`
	ContractDate    string                          `json:"contract_date"`
}

type UpdateUserDataAdminEventData struct {
	ServiceId string `json:"service_id"`
	UserId    string `json:"user_id"`
	Datetime  string `json:"datetime"`
	Message   string `json:"message"`
}

type UpdateUserDataAdmin struct {
	EventType     string                       `json:"event_type"`
	CorrelationId string                       `json:"correlation_id"`
	TelcoId       interface{}                  `json:"telco_id"`
	UserType      interface{}                  `json:"user_type"`
	AbInfo        UpdateUserDataAdminAbInfo    `json:"ab_info"`
	EventData     UpdateUserDataAdminEventData `json:"event_data"`
}
