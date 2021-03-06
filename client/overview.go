package client

import (
	"encoding/json"

	"github.com/gkumar7/goflink"
	"github.com/gkumar7/goflink/http"
)

var _ goflink.Overview = (*Overview)(nil)

type Overview struct {
	HTTPClient *http.Client
}

func (o *Overview) Get() (or *goflink.OverviewResult, err error) {
	d, err := o.HTTPClient.GetBody(&http.Pair{Key: "overview"})
	if err != nil {
		return
	}

	err = json.Unmarshal(d, &or)
	return
}
