package goflink

import "net/http"

type HTTPClient interface {
	Get(string) (*http.Response, error)
}
