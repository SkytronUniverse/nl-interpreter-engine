package web

import (
	"net/http"
)

//HTTPClient interface
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//Client is the web client for making http requests
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
