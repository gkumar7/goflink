package client

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var (
	Props = props{}
)

type props struct {
	FlinkURL string `yaml:"flinkURL"`
}

type client struct {
	Config   *config
	Overview *overview
}

func NewClient(propsFile string) (c *client, err error) {
	f, err := os.Open(propsFile)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &Props)
	if err != nil {
		return
	}

	return &client{
		Config:   new(config),
		Overview: new(overview),
	}, nil
}
