package webhook

type Response struct {
	ID            string `json:"id"`
	ApplicationID string `json:"application_id"`
	Url           string `json:"url"`
	Topic         string `json:"topic"`
	CompanyKey    string `json:"company_key"`
	Active        bool   `json:"active"`
	CreatedAt     string `json:"created_at"`
	Delete        *bool  `json:"delete,omitempty"`
}
