package models

type Additional struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ImId struct {
	ServiceName string `json:"service_name"`
	ServiceId   string `json:"service_id"`
}
