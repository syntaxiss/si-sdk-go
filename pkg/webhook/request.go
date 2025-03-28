package webhook

type Request struct {
	ApplicationID string `json:"application_id"`
	Topic         string `json:"topic"`
	URL           string `json:"url"`
	Active        *bool  `json:"active,omitempty"`
}
