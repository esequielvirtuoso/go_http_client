package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/esequielvirtuoso/go_http_client/gohttp"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start tests for package examples")

	// Tell HTTP library to mock any request comming from test cases
	gohttp.StartMockServer()

	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		assert.Nil(t, endpoints, "no endpoints was expected")
		assert.NotNil(t, err, "an error was expected")
		assert.EqualValues(t, "timeout getting github endpoints", err.Error(), "invalid error message")
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 1}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		assert.Nil(t, endpoints, "no endpoints was expected")
		assert.NotNil(t, err, "an error was expected")
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		assert.NotNil(t, endpoints, "endpoints were expected and we got nil")
		assert.Nil(t, err, "no error was expected")
		assert.EqualValues(t, "https://api.github.com/user", endpoints.CurrentUserUrl, "invalid current user url")
	})

}
