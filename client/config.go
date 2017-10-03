package client

import (
	"encoding/json"

	"github.com/gkumar7/goflink/http"

	"github.com/gkumar7/goflink"
)

var _ goflink.Config = (*Config)(nil)

type Config struct {
	HTTPClient *http.Client
}

func (c *Config) Get() (cr *goflink.ConfigResult, err error) {
	d, err := c.HTTPClient.GetBody(&http.Pair{Key: "config"})
	if err != nil {
		return
	}

	err = json.Unmarshal(d, &cr)
	return
}
