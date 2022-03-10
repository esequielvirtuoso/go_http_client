package gomime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	t.Run("ContentTypeHeader", func(t *testing.T) {
		assert.EqualValues(t, "Content-Type", HeaderContentType, "HeaderContentType should be Content-Type")
	})

	t.Run("UserAgentHeader", func(t *testing.T) {
		assert.EqualValues(t, "User-Agent", HeaderUserAgent, "HeaderUserAgent should be User-Agent")
	})

	t.Run("ContentTypeJson", func(t *testing.T) {
		assert.EqualValues(t, "application/json", ContentTypeJson, "ContentTypeJson should be application/json")
	})

	t.Run("ContentTypeXml", func(t *testing.T) {
		assert.EqualValues(t, "application/xml", ContentTypeXml, "ContentTypeXml should be application/xml")
	})

	t.Run("ContentTypeOctetStream", func(t *testing.T) {
		assert.EqualValues(t, "application/octet-stream", ContentTypeOctetStream, "ContentTypeOctetStream should be application/octet-stream")
	})
}
