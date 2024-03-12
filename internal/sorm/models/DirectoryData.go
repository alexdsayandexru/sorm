package models

type Service struct {
	ServiceId   int32  `json:"service_id"`
	Description string `json:"description"`
	BeginTime   string `json:"begin_time"`
	EndTime     string `json:"end_time"`
}

type DirectoryData struct {
	Datetime      string    `json:"datetime"`
	EventType     string    `json:"event_type"`
	CorrelationId string    `json:"correlation_id"`
	OriServices   []Service `json:"oriServices"`
}

func (target *DirectoryData) GetRules() ValidationRules {
	validationRules := map[string]func() (bool, error){
		"correlation_id": func() (bool, error) {
			return ValidateCorrelationId(target.CorrelationId)
		},
		"datetime": func() (bool, error) {
			return ValidateRequiredDatetime(target.Datetime)
		},
		"oriServices": func() (bool, error) {
			return ValidateOriServices(target.OriServices)
		},
	}
	return validationRules
}
