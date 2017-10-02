package goflink

import "net/http"

// HTTPClient provides http methods used by goflink to access
// Flink's rest api
type HTTPClient interface {
	Get(string) (*http.Response, error)
}
