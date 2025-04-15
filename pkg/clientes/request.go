package clientes

type CustomerType string

const (
	Customer CustomerType = "Customer"
	Supplier CustomerType = "Supplier"
	Other    CustomerType = "Other"
)

type PersonType string

const (
	Person  PersonType = "Person"
	Company PersonType = "Company"
)

type FiscalResponsibility struct {
	Code string `json:"code"`
	Name string `json:"name,omitempty"`
}

type City struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name,omitempty"`
	StateCode   string `json:"state_code"`
	StateName   string `json:"state_name,omitempty"`
	CityCode    string `json:"city_code"`
	CityName    string `json:"city_name,omitempty"`
}

type Address struct {
	Address    string `json:"address"`
	City       City   `json:"city"`
	PostalCode string `json:"postal_code,omitempty"`
}

type Phone struct {
	Indicative string `json:"indicative,omitempty"`
	Number     string `json:"number"`
	Extension  string `json:"extension,omitempty"`
}

type Contact struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     *Phone `json:"phone,omitempty"`
}

type RelatedUsers struct {
	SellerID    int `json:"seller_id,omitempty"`
	CollectorID int `json:"collector_id,omitempty"`
}

type Request struct {
	Type                   CustomerType           `json:"type,omitempty"`
	PersonType             PersonType             `json:"person_type"`
	IDType                 string                 `json:"id_type"`
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
	Comments               string                 `json:"comments,omitempty"`
	RelatedUsers           *RelatedUsers          `json:"related_users,omitempty"`
}
