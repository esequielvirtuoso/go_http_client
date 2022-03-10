package gohttp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")

	builder := clientBuilder{headers: commonHeaders}
	client := httpClient{builder: &builder}

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	fullHeaders := client.getRequestHeaders(requestHeaders)
	// Validation
	assert.EqualValues(t, 3, len(fullHeaders))
	assert.EqualValues(t, "ABC-123", fullHeaders.Get("X-Request-Id"))
	assert.EqualValues(t, "cool-http-client", fullHeaders.Get("User-Agent"))
	assert.EqualValues(t, "application/json", fullHeaders.Get("Content-Type"))
}
