package tsl

import (
	"net/http"
)

// Client for TSL queries
type Client struct {
	endpoint   string
	token      string
	httpClient *http.Client
}

// NewClient return a new TSL client
// httpClient can be nil, default one will be used
func NewClient(endpoint, token string, c *http.Client) *Client {
	cl := &Client{
		endpoint:   endpoint,
		token:      token,
		httpClient: http.DefaultClient,
	}

	if c != nil {
		cl.httpClient = c
	}

	return cl
}

// NewQuery start a new TSL query
func NewQuery(url, token string, c *http.Client) *Query {
	q := &Query{
		endpoint:   url,
		token:      token,
		raw:        "",
		httpClient: http.DefaultClient,
	}

	if c != nil {
		q.httpClient = c
	}

	return q
}

// NewQuery start a new empty query
func (c *Client) NewQuery() *Query {
	return NewQuery(c.endpoint, c.token, c.httpClient)
}
