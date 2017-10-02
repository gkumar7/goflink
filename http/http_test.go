package http_test

import (
	netHTTP "net/http"
	"testing"

	"github.com/gkumar7/goflink/http"
)

var (
	endpointCaptor = new(EndpointCaptor)
	client         = &http.Client{
		HTTP:     endpointCaptor,
		FlinkURL: "localhost",
	}
)

type EndpointCaptor struct {
	val string
}

func (ec *EndpointCaptor) Get(endpoint string) (*netHTTP.Response, error) {
	ec.val = endpoint
	return nil, nil
}

func TestGet(t *testing.T) {
	t.Run("Pair with key and value", func(t *testing.T) {
		client.Get(&http.Pair{"jobs", "1234"})
		expected := "localhost/jobs/1234"
		actual := endpointCaptor.val
		if expected != actual {
			t.Fatalf("Expected %s, but actual was %s\n", expected, actual)
		}
	})

	t.Run("Pair with only key", func(t *testing.T) {
		client.Get(&http.Pair{Key: "jobs"})
		expected := "localhost/jobs"
		actual := endpointCaptor.val
		if expected != actual {
			t.Fatalf("Expected %s, but actual was %s\n", expected, actual)
		}
	})

	t.Run("Multiple pairs", func(t *testing.T) {
		client.Get(
			&http.Pair{"jobs", "1234"},
			&http.Pair{"vertices", "1234"},
			&http.Pair{Key: "taskmanagers"},
		)
		expected := "localhost/jobs/1234/vertices/1234/taskmanagers"
		actual := endpointCaptor.val
		if expected != actual {
			t.Fatalf("Expected %s, but actual was %s\n", expected, actual)
		}
	})
}
