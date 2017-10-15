package goflink

// Config provides methods for /config route
type Config interface {
	Get() (*ConfigResult, error)
}

type ConfigResult struct {
	RefreshInterval int    `json:"refresh-interval"`
	TimeZoneOffset  int    `json:"timezone-offset"`
	TimeZoneName    string `json:"timezone-name"`
	FlinkVersion    string `json:"flink-version"`
	FlinkRevision   string `json:"flink-revision"`
}

type Overview interface {
	Get() (*OverviewResult, error)
}

type OverviewResult struct {
	TaskManagers int `json:"taskmanagers"`
	SlotsResult
	JobCountsResult
}

type SlotsResult struct {
	Total     int `json:"slots-total"`
	Available int `json:"slots-available"`
}

type Jobs interface {
	Get() (*JobIdsResult, error)
}

type JobCountsResult struct {
	Running   int `json:"jobs-running"`
	Finished  int `json:"jobs-finished"`
	Cancelled int `json:"jobs-cancelled"`
	Failed    int `json:"jobs-failed"`
}

type JobIdsResult struct {
	Running   []string `json:"jobs-running"`
	Finished  []string `json:"jobs-finished"`
	Cancelled []string `json:"jobs-cancelled"`
	Failed    []string `json:"jobs-failed"`
}
