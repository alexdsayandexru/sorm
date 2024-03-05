package models

type DeleteUserAccountAdminAdditional []struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeleteUserAccountAdminImId struct {
	ServiceName string `json:"service_name"`
	ServiceId   string `json:"service_id"`
}

type DeleteUserAccountAdminAbInfo struct {
	UserId          string                             `json:"user_id"`
	Msisdn          []string                           `json:"msisdn"`
	Email           []string                           `json:"email"`
	Nick            string                             `json:"nick"`
	BirthDate       string                             `json:"birth_date"`
	Name            string                             `json:"name"`
	Family          string                             `json:"family"`
	Initial         string                             `json:"initial"`
	Address         string                             `json:"address"`
	DatetimeReg     string                             `json:"datetime_reg"`
	DatetimeUpdated string                             `json:"datetime_updated"`
	ServiceId       int                                `json:"service_id"`
	ImId            []DeleteUserAccountAdminImId       `json:"im_id"`
	Additional      []DeleteUserAccountAdminAdditional `json:"additional"`
	DatetimeUnreg   string                             `json:"datetime_unreg"`
	ContractDate    string                             `json:"contract_date"`
}

type DeleteUserAccountAdminEventData struct {
	Datetime  string `json:"datetime"`
	ServiceId string `json:"service_id"`
	UserId    string `json:"user_id"`
}

type DeleteUserAccountAdmin struct {
	EventType     string                          `json:"event_type"`
	CorrelationId string                          `json:"correlation_id"`
	TelcoId       int                             `json:"telco_id"`
	UserType      int                             `json:"user_type"`
	AbInfo        DeleteUserAccountAdminAbInfo    `json:"ab_info"`
	EventData     DeleteUserAccountAdminEventData `json:"event_data"`
}
