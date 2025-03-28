package facturaventa

// TaxResponse representa un impuesto asociado a un item en la respuesta.
type TaxType string

const (
	IVA            TaxType = "IVA"
	Retefuente     TaxType = "Retefuente"
	ReteIVA        TaxType = "ReteIVA"
	ReteICA        TaxType = "ReteICA"
	Impoconsumo    TaxType = "Impoconsumo"
	AdValorem      TaxType = "AdValorem"
	Autorretencion TaxType = "Autorretencion"
)

// Response representa la estructura de respuesta principal de la factura de venta.
type Response struct {
	ID               string                    `json:"id"`
	Document         DocResponse               `json:"document"`
	Number           int                       `json:"number"`
	Name             string                    `json:"name"`
	Date             string                    `json:"date"`
	Customer         CustResponse              `json:"customer"`
	CostCenter       *int                      `json:"cost_center,omitempty"` // Puede ser nulo
	Currency         Currency                  `json:"currency"`
	Total            float64                   `json:"total"`
	Balance          float64                   `json:"balance"`
	Seller           int                       `json:"seller"`
	Stamp            StampResponse             `json:"stamp"`
	Mail             MailResponse              `json:"mail"`
	Observations     *string                   `json:"observations,omitempty"`
	Items            []ItemResponse            `json:"items"`
	Payments         []PaymentResponse         `json:"payments"`
	PublicURL        string                    `json:"public_url"`
	GlobalDiscounts  []GlobalDiscountResponse  `json:"globaldiscounts"`
	AdditionalFields *AdditionalFieldsResponse `json:"additional_fields,omitempty"` //Puede ser nulo
	Metadata         MetadataResponse          `json:"metadata"`
}

// DocResponse representa el tipo de documento en la respuesta.
type DocResponse struct {
	ID int `json:"id"`
}

// CustResponse representa la información del cliente en la respuesta.
type CustResponse struct {
	ID             *string `json:"id,omitempty"` // Puede ser nulo
	Identification string  `json:"identification"`
	BranchOffice   *int    `json:"branch_office,omitempty"` // Puede ser nulo
}

// Currency ya está definido en el request, reutilizar.

// StampResponse representa el estado del envío de la factura electrónica en la respuesta.
type StampResponse struct {
	Status       string `json:"status"`
	Cufe         string `json:"cufe"`
	Observations string `json:"observations"`
	Errors       string `json:"errors"`
}

// MailResponse representa el estado del envío de la factura por correo electrónico en la respuesta.
type MailResponse struct {
	Status       string `json:"status"`
	Observations string `json:"observations"` // Nota: Hay dos campos 'observations' en el JSON de ejemplo para 'mail'. Asumo que es un error y que debería ser uno solo. Si son distintos, renombrar o añadir un segundo campo si es necesario.
}

type MailRequest struct {
	MailTo string `json:"mail_to"`
	CopyTo string `json:"copy_to"`
}

// ItemResponse representa un producto o servicio en la respuesta de la factura.
type ItemResponse struct {
	ID          string            `json:"id"`
	Code        string            `json:"code"`
	Description string            `json:"description"`
	Quantity    float64           `json:"quantity"`
	Price       float64           `json:"price"`
	Discount    *DiscountResponse `json:"discount,omitempty"` // Puede ser nulo
	Taxes       []TaxResponse     `json:"taxes,omitempty"`    // Puede ser nulo
	Total       float64           `json:"total"`
}

// DiscountResponse representa el descuento de un item en la respuesta.
type DiscountResponse struct {
	Percentage float64 `json:"percentage"`
	Value      float64 `json:"value"`
}

type TaxResponse struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Type       TaxType `json:"type"` // Usar TaxType para enum
	Percentage float64 `json:"percentage"`
	Value      float64 `json:"value"`
}

// PaymentResponse representa una forma de pago en la respuesta.
type PaymentResponse struct {
	ID      *int    `json:"id,omitempty"` // Puede ser nulo
	Name    string  `json:"name"`
	Value   float64 `json:"value"`
	DueDate string  `json:"due_date"`
}

// GlobalDiscountResponse representa un descuento global en la respuesta de la factura.
type GlobalDiscountResponse struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Percentage float64 `json:"percentage"`
	Value      float64 `json:"value"`
}

// AdditionalFieldsResponse representa campos adicionales en la respuesta.
// Similar al Request, se deja vacío a falta de especificación de campos concretos.
type AdditionalFieldsResponse struct {
	//Campos adicionales aquí si se conocen
}

// MetadataResponse representa la metadata de la respuesta.
type MetadataResponse struct {
	Created     string  `json:"created"`
	LastUpdated *string `json:"last_updated,omitempty"` // Puede ser nulo
}
