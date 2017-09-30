package http_test

import (
	"goflink/http"
	stdHTTP "net/http"
	"testing"
)

var (
	mockHTTPClient = new(MockHTTPClient)
	client         = &http.Client{
		HTTP:     mockHTTPClient,
		FlinkURL: "localhost",
	}
)

type MockHTTPClient struct {
	endpointCapture string
}

func (m *MockHTTPClient) Get(endpoint string) (*stdHTTP.Response, error) {
	m.endpointCapture = endpoint
	return nil, nil
}

func TestGet(t *testing.T) {
	client.Get(&http.Pair{"jobs", "1234"})
	expected := "localhost/jobs/1234"
	actual := mockHTTPClient.endpointCapture
	if expected != actual {
		t.Fatalf("Expected %s, but actual was %s\n", expected, actual)
	}
}
