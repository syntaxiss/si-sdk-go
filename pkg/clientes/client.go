package clientes

import (
	"context"
	"net/http"

	"github.com/syntaxiss/si-sdk-go/pkg/configuracion"
	"github.com/syntaxiss/si-sdk-go/pkg/internal/httpclient"
)

const (
	urlBase           = "https://api.siigo.com"
	urlCrearCliente   = urlBase + "/v1/customers"
	urlObtenerCliente = urlBase + "/v1/customers/" + "{id}"
)

type Client interface {
	CrearCliente(ctx context.Context, request Request) (*Response, error)
	ObtenerCliente(ctx context.Context, id string) (*Response, error)
	BuscarCliente(ctx context.Context, request SearchRequest) (*SearchResponse, error)
	ActualizarCliente(ctx context.Context, request Request, id string) (*Response, error)
}

type client struct {
	configuracion *configuracion.Configuracion
}

func NewClient(c *configuracion.Configuracion) Client {
	return &client{
		configuracion: c,
	}
}

func (c *client) CrearCliente(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlCrearCliente,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ObtenerCliente(ctx context.Context, id string) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlObtenerCliente,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) BuscarCliente(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		Method:      http.MethodGet,
		URL:         urlCrearCliente,
		QueryParams: queryParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ActualizarCliente(ctx context.Context, request Request, id string) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPut,
		URL:        urlObtenerCliente,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
