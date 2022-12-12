package http_helper

import (
	"net/http"
)

// please finish the annotation of the part
type FakeResponse struct {
	headers http.Header
	body    []byte
	status  int
}

func NewFakeResponse() *FakeResponse {
	return &FakeResponse{
		headers: make(http.Header),
	}
}

func (r *FakeResponse) Header() http.Header {
	return r.headers
}

func (r *FakeResponse) Write(body []byte) (int, error) {
	r.body = body
	return len(body), nil
}

func (r *FakeResponse) WriteHeader(status int) {
	r.status = status
}

func (r *FakeResponse) GetBody() []byte {
	return r.body
}
