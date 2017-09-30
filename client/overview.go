package client

import "github.com/gkumar7/goflink"

type overview struct {
}

func (o *overview) Get() *goflink.Summary {
	return nil
}
