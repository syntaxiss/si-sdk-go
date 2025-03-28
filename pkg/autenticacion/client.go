package autenticacion

import (
	"context"
	"net/http"

	"github.com/syntaxiss/si-sdk-go/pkg/configuracion"
	"github.com/syntaxiss/si-sdk-go/pkg/internal/httpclient"
)

const (
	urlBase         = "https://api.siigo.com"
	urlGenerarToken = urlBase + "/auth"
)

type Client interface {
	GenerarToken(ctx context.Context, request Request) (*Response, error)
}

type client struct {
	configuracion *configuracion.Configuracion
}

func NewClient(c *configuracion.Configuracion) Client {
	return &client{
		configuracion: c,
	}
}

func (c *client) GenerarToken(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlGenerarToken,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
