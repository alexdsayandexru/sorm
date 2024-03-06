package models

type DeleteAccountImId struct {
	ServiceName string `json:"service_name"`
	ServiceId   string `json:"service_id"`
}

type DeleteAccountAdditional struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeleteAccountAbInfo struct {
	UserId          string                    `json:"user_id"`
	Nick            string                    `json:"nick"`
	BirthDate       string                    `json:"birth_date"`
	Name            string                    `json:"name"`
	Family          string                    `json:"family"`
	Initial         string                    `json:"initial"`
	Msisdn          []string                  `json:"msisdn"`
	Email           []string                  `json:"email"`
	Address         string                    `json:"address"`
	DatetimeReg     string                    `json:"datetime_reg"`
	DatetimeUpdated string                    `json:"datetime_updated"`
	ServiceId       int                       `json:"service_id"`
	ImId            []DeleteAccountImId       `json:"im_id"`
	Additional      []DeleteAccountAdditional `json:"additional"`
	DatetimeUnreg   int                       `json:"datetime_unreg"`
	ContractDate    string                    `json:"contract_date"`
}

type DeleteAccount struct {
	EventType     string              `json:"event_type"`
	CorrelationId string              `json:"correlation_id"`
	TelcoId       int                 `json:"telco_id"`
	UserType      int                 `json:"user_type"`
	AbInfo        DeleteAccountAbInfo `json:"ab_info"`
}
