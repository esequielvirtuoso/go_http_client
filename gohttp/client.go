package gohttp

import (
	"net/http"
	"sync"

	"github.com/esequielvirtuoso/go_http_client/core"
)

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

type Client interface {
	Get(string, ...http.Header) (*core.Response, error)
	Post(string, interface{}, ...http.Header) (*core.Response, error)
	Put(string, interface{}, ...http.Header) (*core.Response, error)
	Patch(string, interface{}, ...http.Header) (*core.Response, error)
	Delete(string, ...http.Header) (*core.Response, error)
	Options(string, ...http.Header) (*core.Response, error)
}

func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), body)
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}

func (c *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return nil, nil
}
