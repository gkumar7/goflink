package client

import (
	"io/ioutil"
	"os"

	"github.com/gkumar7/goflink"
	"github.com/gkumar7/goflink/http"

	yaml "gopkg.in/yaml.v2"
)

type props struct {
	FlinkURL string `yaml:"flinkURL"`
}

// A Client is structured such that each of its members correspond to
// a single base route of the Flink REST API
type Client struct {
	Config   goflink.Config
	Overview goflink.Overview
}

// NewClient creates a goflink REST Client
func NewClient(propsFile string) (c *Client, err error) {
	f, err := os.Open(propsFile)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	var p props
	err = yaml.Unmarshal(data, &p)
	if err != nil {
		return
	}

	HTTPClient := http.NewClient(p.FlinkURL)
	return &Client{
		Config:   &Config{HTTPClient},
		Overview: &Overview{HTTPClient},
	}, nil
}
