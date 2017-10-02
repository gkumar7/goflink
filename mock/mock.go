package mock

import (
	"bytes"
	"net/http"
)

type HTTPMock struct {
	ReturnData string
}

func (m *HTTPMock) Get(endpoint string) (*http.Response, error) {
	buf := new(CloseableBuffer)
	resp := new(http.Response)
	buf.WriteString(m.ReturnData)
	resp.Body = buf
	return resp, nil
}

type CloseableBuffer struct {
	bytes.Buffer
}

func (cb *CloseableBuffer) Close() error {
	return nil
}
