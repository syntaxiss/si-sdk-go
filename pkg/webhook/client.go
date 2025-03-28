package webhook

import (
	"context"
	"net/http"

	"github.com/syntaxiss/si-sdk-go/pkg/configuracion"
	"github.com/syntaxiss/si-sdk-go/pkg/internal/httpclient"
)

const (
	urlBase              = "https://api.siigo.com/v1"
	urlSubscribirWebhook = urlBase + "/webhooks"
	urlEditarWebhook     = urlBase + "//webhooks"
	urlFindWebhook       = urlBase + "/ webhooks"
	urlEliminarWebhook   = urlBase + "/webhooks/" + "{id}"
)

type Client interface {
	SubscribeWebhook(ctx context.Context, Request Request) (*Response, error)
	EditWebhook(ctx context.Context, Request Request) (*Response, error)
	FindWebhook(ctx context.Context, Request Request) (*SearchResponse, error)
	DeleteWebhook(ctx context.Context, id string) (*Response, error)
}

type client struct {
	configuracion *configuracion.Configuracion
}

func NewClient(c *configuracion.Configuracion) Client {
	return &client{
		configuracion: c,
	}
}

func (c *client) SubscribeWebhook(ctx context.Context, Request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   Request,
		Method: http.MethodPost,
		URL:    urlSubscribirWebhook,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) EditWebhook(ctx context.Context, Request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   Request,
		Method: http.MethodPut,
		URL:    urlEditarWebhook,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) FindWebhook(ctx context.Context, Request Request) (*SearchResponse, error) {
	requestData := httpclient.RequestData{
		Body:   Request,
		Method: http.MethodGet,
		URL:    urlFindWebhook,
	}

	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) DeleteWebhook(ctx context.Context, id string) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodDelete,
		URL:        urlEliminarWebhook,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
