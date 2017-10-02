package client_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gkumar7/goflink"
	"github.com/gkumar7/goflink/client"
	"github.com/gkumar7/goflink/http"
	"github.com/gkumar7/goflink/mock"
)

func TestNewClient(t *testing.T) {
	client, err := client.NewClient("props_test.yaml")
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if client == nil {
		t.Fatal("Client unexpectedly nil")
	}
}

func TestGet(t *testing.T) {
	expected := &goflink.ConfigResult{
		RefreshInterval: 3000,
		TimeZoneOffset:  -21600000,
		TimeZoneName:    "Central Standard Time",
		FlinkVersion:    "1.3.2",
		FlinkRevision:   "0399bee @ 03.08.2017 @ 10:23:11 UTC",
	}
	mock := &mock.HTTPMock{
		ReturnData: fmt.Sprintf(
			`{"refresh-interval": %d,
                          "timezone-offset": %d,
                          "timezone-name": "%s",
                          "flink-version": "%s",
                          "flink-revision": "%s"}`,
			expected.RefreshInterval,
			expected.TimeZoneOffset,
			expected.TimeZoneName,
			expected.FlinkVersion,
			expected.FlinkRevision,
		),
	}

	httpClient := &http.Client{HTTP: mock}
	client := &client.Client{
		Config:   &client.Config{httpClient},
		Overview: &client.Overview{httpClient},
	}

	actual, err := client.Config.Get()
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
}
