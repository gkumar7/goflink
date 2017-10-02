package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gkumar7/goflink/http"

	"github.com/gkumar7/goflink"
)

var _ goflink.Config = (*Config)(nil)

type Config struct {
	HTTPClient *http.Client
}

func (c *Config) Get() (cr *goflink.ConfigResult, err error) {
	resp, err := c.HTTPClient.Get(&http.Pair{Key: "config"})
	if err != nil {
		return
	}

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(d, &cr)
	if err != nil {
		return
	}

	return
}
