package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
	baseUrl            string
	client             *http.Client
	userAgent          string
}

type ClientBuilder interface {
	SetHeaders(http.Header) ClientBuilder
	SetConnectionTimeout(time.Duration) ClientBuilder
	SetRequestTimeout(time.Duration) ClientBuilder
	SetIdleConnections(int) ClientBuilder
	DisableTimeouts(bool) ClientBuilder
	SetHttpClient(*http.Client) ClientBuilder
	SetUserAgent(string) ClientBuilder

	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetRequestTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnections = connections
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(agent string) ClientBuilder {
	c.userAgent = agent
	return c
}
