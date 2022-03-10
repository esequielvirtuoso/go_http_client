package examples

import (
	"net/http"
	"time"

	"github.com/esequielvirtuoso/go_http_client/gohttp"
	"github.com/esequielvirtuoso/go_http_client/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	// setting an http.Client is an optional configuration
	// WARN: it will bypass all the remaining configurations such as the timeouts
	// currentClient := http.Client{}
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetRequestTimeout(3 * time.Second).
		SetUserAgent("test-agent-client").
		// SetHttpClient(&currentClient).
		Build()

	return client
}
