package client

import (
	"encoding/json"

	"github.com/gkumar7/goflink"
	"github.com/gkumar7/goflink/http"
)

var _ goflink.Jobs = (*Jobs)(nil)

type Jobs struct {
	HTTPClient *http.Client
}

func (j *Jobs) Get() (jr *goflink.JobIdsResult, err error) {
	d, err := j.HTTPClient.GetBody(&http.Pair{Key: "jobs"})
	if err != nil {
		return
	}

	err = json.Unmarshal(d, &jr)
	return
}
