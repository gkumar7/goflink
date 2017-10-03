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

func (o *Jobs) Get() (or *goflink.JobsResult, err error) {
	d, err := o.HTTPClient.GetBody(&http.Pair{Key: "jobs"})
	if err != nil {
		return
	}

	err = json.Unmarshal(d, &or)
	return
}
