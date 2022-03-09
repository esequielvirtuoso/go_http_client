package examples

import (
	"time"

	"github.com/esequielvirtuoso/go_http_client/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetRequestTimeout(3 * time.Second).
		Build()

	return client
}
