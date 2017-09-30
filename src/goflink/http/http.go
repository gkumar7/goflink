package http

import (
	"bytes"
	"fmt"
	"goflink"
	"net/http"
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
	_, err = buf.WriteString(p.Key)
	_, err = buf.WriteRune('/')
	_, err = buf.WriteString(p.Value)
	return
}

func (c *Client) Get(pairs ...*Pair) (resp *http.Response, err error) {
	var endpoint bytes.Buffer
	for _, p := range pairs {
		writePair(p, &endpoint)
	}
	return c.HTTP.Get(fmt.Sprintf("%s/%s", c.FlinkURL, endpoint.String()))
}

func NewClient(flinkURL string) *Client {
	return &Client{
		HTTP:     new(http.Client),
		FlinkURL: flinkURL,
	}
}
