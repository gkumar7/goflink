package client_test

import (
	"encoding/json"
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

func TestConfigGet(t *testing.T) {
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
		Config: &client.Config{httpClient},
	}

	actual, err := client.Config.Get()
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
}

func TestOverviewGet(t *testing.T) {
	expected := &goflink.OverviewResult{
		TaskManagers:    17,
		SlotsResult:     goflink.SlotsResult{Total: 68, Available: 68},
		JobCountsResult: goflink.JobCountsResult{Running: 0, Finished: 3, Cancelled: 1, Failed: 0},
	}
	mock := &mock.HTTPMock{
		ReturnData: fmt.Sprintf(
			`{"taskmanagers": %d,
                          "slots-total": %d,
                          "slots-available": %d,
                          "jobs-running": %d,
                          "jobs-finished": %d,
                          "jobs-cancelled": %d,
                          "jobs-failed": %d}`,
			expected.TaskManagers,
			expected.SlotsResult.Total,
			expected.SlotsResult.Available,
			expected.JobCountsResult.Running,
			expected.JobCountsResult.Finished,
			expected.JobCountsResult.Cancelled,
			expected.JobCountsResult.Failed,
		),
	}

	httpClient := &http.Client{HTTP: mock}
	client := &client.Client{
		Overview: &client.Overview{httpClient},
	}

	actual, err := client.Overview.Get()
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
}

func getJson(lst []string) string {
	b, _ := json.Marshal(lst)
	return string(b)
}

func TestJobsGet(t *testing.T) {
	expected := &goflink.JobIdsResult{
		Running:   []string{"a", "b"},
		Finished:  []string{"a", "b", "c"},
		Cancelled: []string{"a"},
		Failed:    []string{},
	}
	mock := &mock.HTTPMock{
		ReturnData: fmt.Sprintf(
			`{"jobs-running": %s,
                          "jobs-finished": %s,
                          "jobs-cancelled": %s,
                          "jobs-failed": %s
                         }`,
			getJson(expected.Running),
			getJson(expected.Finished),
			getJson(expected.Cancelled),
			getJson(expected.Failed),
		),
	}

	httpClient := &http.Client{HTTP: mock}
	client := &client.Client{
		Jobs: &client.Jobs{httpClient},
	}

	actual, err := client.Jobs.Get()
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, actual %v", expected, actual)
	}
}
