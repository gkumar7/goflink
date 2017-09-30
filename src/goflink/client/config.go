package client

import "goflink"

var _ goflink.Config = (*config)(nil)

type config struct {
}

func (c *config) Get() *goflink.ConfigResult {
	return nil
}
