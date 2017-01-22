package openuri

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// Client type
type Client struct {
	httpClient *http.Client
}

// ClientOption type
type ClientOption func(*Client) error

// New returns a Client struct
func New(options ...ClientOption) (*Client, error) {
	c := &Client{httpClient: http.DefaultClient}
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Open an io.ReadCloser from a local file or URL
func Open(name string, options ...ClientOption) (io.ReadCloser, error) {
	c, err := New(options...)
	if err != nil {
		return nil, err
	}
	return c.Open(name)
}

func WithHTTPClient(v *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = v
		return nil
	}
}

func (c *Client) Open(name string) (io.ReadCloser, error) {
	if strings.HasPrefix(name, "http://") || strings.HasPrefix(name, "https://") {
		resp, err := c.httpClient.Get(name)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}
	return os.Open(name)
}
