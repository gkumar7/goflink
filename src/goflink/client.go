package goflink

type ConfigResult struct {
	RefreshInterval int
	TimeZoneOffset  int
	TimeZoneName    int
	FlinkVersion    int
	FlinkRevision   int
}

type Config interface {
	Get() *ConfigResult
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
