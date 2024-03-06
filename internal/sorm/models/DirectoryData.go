package models

type OriServices struct {
	ServiceId   string `json:"service_id"`
	Description string `json:"description"`
	BeginTime   string `json:"begin_time"`
	EndTime     string `json:"end_time"`
}

type DirectoryData struct {
	Datetime      string        `json:"datetime"`
	EventType     string        `json:"event_type"`
	CorrelationId string        `json:"correlation_id"`
	OriServices   []OriServices `json:"oriServices"`
}
