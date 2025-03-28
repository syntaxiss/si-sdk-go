package configuracion

import (
	"fmt"

	"github.com/syntaxiss/si-sdk-go/pkg/internal/defaultrequester"
	"github.com/syntaxiss/si-sdk-go/pkg/requester"
)

type Configuracion struct {
	Requester requester.Requester

	Authorization string
}

func New(authorization string, opts ...Option) (*Configuracion, error) {
	cfg := &Configuracion{
		Authorization: authorization,
		Requester:     defaultrequester.New(),
	}

	// Apply all the functional options to configure the client.
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, fmt.Errorf("fail to build config: %w", err)
		}
	}

	return cfg, nil
}
