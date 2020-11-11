package web

import (
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	HTTPClient HTTPClient
}

//NewWebClient creates a new http client
func NewWebClient(h HTTPClient) *Client {

	w := &Client{
		HTTPClient: http.DefaultClient,
	}

	if h != nil {
		w.HTTPClient = h
	}
	return w
}
