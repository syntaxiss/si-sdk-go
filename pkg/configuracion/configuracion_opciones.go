package configuracion

import (
	"fmt"

	"github.com/syntaxiss/si-sdk-go/pkg/requester"
)

type Option func(*Configuracion) error

// WithHTTPClient allow to do requests using received http client.
func WithHTTPClient(r requester.Requester) Option {
	return func(c *Configuracion) error {
		if r == nil {
			return fmt.Errorf("received http client can't be nil")
		}
		c.Requester = r
		return nil
	}
}
