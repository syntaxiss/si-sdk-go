package clientes

type IDType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Metadata struct {
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated,omitempty"`
}

type Response struct {
	ID                     string                 `json:"id"`
	Type                   CustomerType           `json:"type,omitempty"` // Assuming 'type' refers to the customer type
	PersonType             PersonType             `json:"person_type"`
	IDType                 IDType                 `json:"id_type"`
	Identification         string                 `json:"identification"`
	CheckDigit             string                 `json:"check_digit,omitempty"`
	Name                   []string               `json:"name"`
	CommercialName         string                 `json:"commercial_name,omitempty"`
	BranchOffice           int                    `json:"branch_office,omitempty"`
	Active                 bool                   `json:"active,omitempty"`
	VatResponsible         bool                   `json:"vat_responsible,omitempty"`
	FiscalResponsibilities []FiscalResponsibility `json:"fiscal_responsibilities,omitempty"`
	Address                Address                `json:"address"`
	Phones                 []Phone                `json:"phones"`
	Contacts               []Contact              `json:"contacts"`
	RelatedUsers           *RelatedUsers          `json:"related_users,omitempty"`
	Metadata               *Metadata              `json:"metadata,omitempty"`
}
