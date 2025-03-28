package httpclient

import (
	"fmt"
	"io"
	"net/http"

	"github.com/syntaxiss/si-sdk-go/pkg/requester"
	"github.com/syntaxiss/si-sdk-go/pkg/siierror"
)

func Send(requester requester.Requester, req *http.Request) ([]byte, error) {
	res, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %w", err)
	}

	defer func() { _ = res.Body.Close() }()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &siierror.ResponseError{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode > 399 {
		return nil, &siierror.ResponseError{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return response, nil
}
