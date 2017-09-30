package client

import "goflink"

type overview struct {
}

func (o *overview) Get() *goflink.Summary {
	return nil
}
