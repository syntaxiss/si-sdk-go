package facturaventa

import (
	"context"
	"net/http"

	"github.com/syntaxiss/si-sdk-go/pkg/configuracion"
	"github.com/syntaxiss/si-sdk-go/pkg/internal/httpclient"
)

const (
	urlBase                = "https://api.siigo.com"
	urlCrearFacturaVenta   = urlBase + "/v1/invoices"
	urlObtenerFacturaVenta = urlBase + "/v1/invoices/" + "{id}"
	urlAnularFacturaVenta  = urlObtenerFacturaVenta + "/annul"
	urlEnviarCorreo        = urlObtenerFacturaVenta + "/mail"
	urlObtenerPdf          = urlObtenerFacturaVenta + "/pdf"
	urlObtenerXml          = urlObtenerFacturaVenta + "/xml"
)

type Client interface {
	CrearFacturaVenta(ctx context.Context, request Request) (*Response, error)
	ObtenerFacturaVenta(ctx context.Context, id string) (*Response, error)
	EliminarFacturaVenta(ctx context.Context, id string) error
	AnularFacturaventa(ctx context.Context, id string) error
	BuscarFacturaVenta(ctx context.Context, request SearchRequest) (*SearchResponse, error)
	EnviarAlCorreo(ctx context.Context, request MailRequest, id string) (*MailResponse, error)
	ErroresDian(ctx context.Context, id string) (*ErrorDianResponse, error)
}

type client struct {
	configuracion *configuracion.Configuracion
}

func NewClient(c *configuracion.Configuracion) Client {
	return &client{
		configuracion: c,
	}
}

func (c *client) CrearFacturaVenta(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlCrearFacturaVenta,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ObtenerFacturaVenta(ctx context.Context, id string) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlObtenerFacturaVenta,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) EliminarFacturaVenta(ctx context.Context, id string) error {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodDelete,
		URL:        urlObtenerFacturaVenta,
		PathParams: pathParams,
	}

	_, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	return err
}

func (c *client) AnularFacturaventa(ctx context.Context, id string) error {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlAnularFacturaVenta,
		PathParams: pathParams,
	}

	_, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	return err
}

func (c *client) BuscarFacturaVenta(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		Method:      http.MethodGet,
		URL:         urlCrearFacturaVenta,
		QueryParams: queryParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) EnviarAlCorreo(ctx context.Context, request MailRequest, id string) (*MailResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPost,
		URL:        urlEnviarCorreo,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*MailResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ObtenerPdf(ctx context.Context, id string) (*DocmentsResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlObtenerPdf,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*DocmentsResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ObtenerXml(ctx context.Context, id string) (*DocmentsResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlObtenerXml,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*DocmentsResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ErroresDian(ctx context.Context, id string) (*ErrorDianResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlObtenerFacturaVenta,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*ErrorDianResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
