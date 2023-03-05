package apize

import "net/http"

// Client is an Apize API client. It handles authentication and HTTP requests to interact with Apize APIs.
type Client interface {
	TextClient
}

type ClientOption func(c *client)

// New creates a new Apize Client instance with an API token and additional options.
func New(token string, opts ...ClientOption) Client {
	c := &client{
		token:      token,
		httpClient: http.DefaultClient,
		baseurl:    "https://api.apize.dev",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type client struct {
	token      string
	httpClient *http.Client
	baseurl    string
}

// WithHttpClient sets the HTTP Client to be used when sending HTTP requests to Apize.
func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

// WithBaseUrl sets the base URL for the Apize client. By default, the base URL is "https://api.apize.dev".
func WithBaseUrl(baseurl string) ClientOption {
	return func(c *client) {
		c.baseurl = baseurl
	}
}
