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
	Get() *Summary
}

type Summary struct {
	TaskManagers int
	Slots
	Jobs
}

type Slots struct {
	Total     int
	Available int
}

type Jobs struct {
	Running   int
	Finished  int
	Cancelled int
	Failed    int
}
