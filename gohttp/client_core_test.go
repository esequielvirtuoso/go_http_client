package gohttp

import (
	"testing"

	"github.com/esequielvirtuoso/go_http_client/gomime"
	"github.com/stretchr/testify/assert"
)

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}
	t.Run("NoBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)

		// Validation
		assert.Nil(t, err)
		assert.Nil(t, body)
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		// Initialization
		requestBody := []string{"one", "two"}

		// Execution
		body, err := client.getRequestBody(gomime.ContentTypeJson, requestBody)

		// Validation
		assert.Nil(t, err)
		assert.NotNil(t, body)
		assert.EqualValues(t, `["one","two"]`, string(body))
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		// Initialization
		requestBody := []string{"one", "two"}

		// Execution
		body, err := client.getRequestBody(gomime.ContentTypeXml, requestBody)

		// Validation
		assert.Nil(t, err)
		assert.NotNil(t, body)
		assert.EqualValues(t, `<string>one</string><string>two</string>`, string(body))
	})

	t.Run("BodyWithDefault", func(t *testing.T) {
		// Initialization
		requestBody := []string{"one", "two"}

		// Execution
		body, err := client.getRequestBody("", requestBody)

		// Validation
		assert.Nil(t, err)
		assert.NotNil(t, body)
		assert.EqualValues(t, `["one","two"]`, string(body))
	})
}
