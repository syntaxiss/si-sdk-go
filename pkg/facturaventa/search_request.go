package facturaventa

import (
	"strings"
)

type SearchRequest struct {
	Filters map[string]string
}

// GetParams creates map to build query parameters. Keys will be changed to lower case.
func (sr *SearchRequest) GetParams() map[string]string {
	params := map[string]string{}
	for k, v := range sr.Filters {
		key := strings.ToLower(k)
		params[key] = v
	}

	return params
}
