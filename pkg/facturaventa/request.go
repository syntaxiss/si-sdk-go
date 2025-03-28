package facturaventa

// Request es la estructura principal que contiene todos los atributos de la solicitud.
type Request struct {
	Document         Document          `json:"document"`
	Date             string            `json:"date"` // formato Fecha de la factura, formato yyyy-MM-dd.
	Customer         Customer          `json:"customer"`
	CostCenter       *int              `json:"cost_center,omitempty"` // Puede ser nulo
	Currency         *Currency         `json:"currency,omitempty"`
	Seller           int               `json:"seller"`
	Stamp            Stamp             `json:"stamp"`
	Mail             Mail              `json:"mail"`
	Observations     *string           `json:"observations,omitempty"` // Puede ser nulo
	Items            []Item            `json:"items"`
	Transport        *Transport        `json:"transport,omitempty"` // Puede ser nulo
	Payments         []Payment         `json:"payments"`
	GlobalDiscounts  []GlobalDiscount  `json:"globaldiscounts"`
	AdditionalFields *AdditionalFields `json:"additional_fields,omitempty"` // Puede ser nulo
}

// Document representa el tipo de documento y su identificador.
type Document struct {
	ID int `json:"id"`
}

// Customer representa la información del cliente.
type Customer struct {
	PersonType     string    `json:"person_type"`
	IDType         string    `json:"id_type"`
	Identification string    `json:"identification"`
	BranchOffice   *int      `json:"branch_office,omitempty"` // Puede ser nulo
	Name           []string  `json:"name"`                    // Usando []string para representar "Razón social o nombres y apellidos del cliente"
	Address        Address   `json:"address"`
	Phones         []Phone   `json:"phones"`
	Contacts       []Contact `json:"contacts"`
}

// Address representa la dirección del cliente.
type Address struct {
	Address    string  `json:"address"`
	City       City    `json:"city"`
	PostalCode *string `json:"postal_code,omitempty"` // Puede ser nulo
}

// City representa la ciudad del cliente.
type City struct {
	CountryCode string  `json:"country_code"`
	CountryName *string `json:"country_name,omitempty"` // Puede ser nulo
	StateCode   string  `json:"state_code"`
	StateName   *string `json:"state_name,omitempty"` // Puede ser nulo
	CityCode    string  `json:"city_code"`
	CityName    *string `json:"city_name,omitempty"` // Puede ser nulo
}

// Phone representa un número de teléfono.
type Phone struct {
	Indicative string  `json:"indicative"`
	Number     string  `json:"number"`
	Extension  *string `json:"extension,omitempty"` // Puede ser nulo
}

// Contact representa un contacto asociado al cliente.
type Contact struct {
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Phone     *ContactPhone `json:"phone,omitempty"` // Puede ser nulo
}

// ContactPhone representa el teléfono de un contacto.
type ContactPhone struct {
	Indicative string  `json:"indicative"`
	Number     string  `json:"number"`
	Extension  *string `json:"extension,omitempty"` // Puede ser nulo
}

// Currency representa la moneda utilizada.
type Currency struct {
	Code         string  `json:"code"`
	ExchangeRate float64 `json:"exchange_rate"`
}

// Stamp representa la información para el envío de factura electrónica.
type Stamp struct {
	Send bool `json:"send"`
}

// Mail representa la información para el envío de la factura por correo electrónico.
type Mail struct {
	Send bool `json:"send"`
}

// Item representa un producto o servicio en la factura.
type Item struct {
	Code        string         `json:"code"`
	Description *string        `json:"description,omitempty"` // Puede ser nulo
	Quantity    float64        `json:"quantity"`
	Price       float64        `json:"price"`
	Discount    *float64       `json:"discount,omitempty"`  // Puede ser nulo
	Taxes       []Tax          `json:"taxes,omitempty"`     // Puede ser nulo
	Transport   *ItemTransport `json:"transport,omitempty"` // Puede ser nulo - Transport dentro de Item para transport de línea
}

// Tax representa un impuesto asociado a un item.
type Tax struct {
	ID int `json:"id"`
}

// ItemTransport representa la información de transporte dentro de un item (línea).
type ItemTransport struct {
	FileNumber          *int     `json:"file_number,omitempty"`          // Puede ser nulo
	ShipmentNumber      *string  `json:"shipment_number,omitempty"`      // Puede ser nulo
	TransportedQuantity *float64 `json:"transported_quantity,omitempty"` // Puede ser nulo
	MeasurementUnit     *string  `json:"measurement_unit,omitempty"`     // Puede ser nulo
	FreightValue        *float64 `json:"freight_value,omitempty"`        // Puede ser nulo
	PurchaseOrder       *string  `json:"purchase_order,omitempty"`       // Puede ser nulo
	ServiceType         *string  `json:"service_type,omitempty"`         // Puede ser nulo
}

// Transport representa la información de transporte a nivel de factura.
// Nota: Hay un struct ItemTransport para transporte dentro de cada Item. Este struct Transport es diferente y a nivel global de la factura.
type Transport struct {
	FileNumber          *int     `json:"file_number,omitempty"`          // Puede ser nulo
	ShipmentNumber      *string  `json:"shipment_number,omitempty"`      // Puede ser nulo
	TransportedQuantity *float64 `json:"transported_quantity,omitempty"` // Puede ser nulo
	MeasurementUnit     *string  `json:"measurement_unit,omitempty"`     // Puede ser nulo
	FreightValue        *float64 `json:"freight_value,omitempty"`        // Puede ser nulo
	PurchaseOrder       *string  `json:"purchase_order,omitempty"`       // Puede ser nulo
	ServiceType         *string  `json:"service_type,omitempty"`         // Puede ser nulo
}

// Payment representa una forma de pago.
type Payment struct {
	ID      int     `json:"id"`
	Value   float64 `json:"value"`
	DueDate *string `json:"due_date,omitempty"` // Puede ser nulo
}

// GlobalDiscount representa un descuento global en la factura.
type GlobalDiscount struct {
	ID         int     `json:"id"`
	Percentage float64 `json:"percentage"`
	Value      float64 `json:"value"`
}

// AdditionalFields representa campos adicionales.  En la descripción no se especifican campos dentro de additional_fields
// Por lo tanto se deja como un struct vacío, pero puede ser modificado si se conocen los campos.
type AdditionalFields struct {
	ExtraInfo string `json:"extra_info"`
}

type DocmentsResponse struct {
	Id     string `json:"id"`
	Base64 string `json:"base64"`
}

type ErrorDianResponse struct {
	Id      string  `json:"id"`
	Errores []Error `json:"errores"`
}

type Error struct {
	Menssage string `json:"message"`
}
