package client

import "github.com/gkumar7/goflink"

var _ goflink.Config = (*config)(nil)

type config struct {
}

func (c *config) Get() *goflink.ConfigResult {
	return nil
}
