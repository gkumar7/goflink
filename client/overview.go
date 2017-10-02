package client

import (
	"github.com/gkumar7/goflink"
	"github.com/gkumar7/goflink/http"
)

type Overview struct {
	HTTPClient *http.Client
}

func (o *Overview) Get() *goflink.Summary {
	return nil
}
