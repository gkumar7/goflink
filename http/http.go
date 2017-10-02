package http

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gkumar7/goflink"
)

type Client struct {
	HTTP     goflink.HTTPClient
	FlinkURL string
}

type Pair struct {
	Key   string
	Value string
}

func writePair(p *Pair, buf *bytes.Buffer) (err error) {
	_, err = buf.WriteRune('/')
	_, err = buf.WriteString(p.Key)

	if len(p.Value) > 0 {
		_, err = buf.WriteRune('/')
		_, err = buf.WriteString(p.Value)
	}

	return
}

func (c *Client) Get(pairs ...*Pair) (resp *http.Response, err error) {
	var endpoint bytes.Buffer
	for _, p := range pairs {
		if err = writePair(p, &endpoint); err != nil {
			return
		}
	}
	return c.HTTP.Get(fmt.Sprintf("%s%s", c.FlinkURL, endpoint.String()))
}

func NewClient(flinkURL string) *Client {
	return &Client{
		HTTP:     new(http.Client),
		FlinkURL: flinkURL,
	}
}
