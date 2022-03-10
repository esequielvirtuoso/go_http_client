package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/esequielvirtuoso/go_http_client/gohttp"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepo(t *testing.T) {
	t.Run("TestTimeoutFromGithub", func(t *testing.T) {
		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","description":"Test repository.","private":true}`,
			Error:       errors.New("timeout from github"),
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "Test repository.",
			Private:     true,
		}

		// Execution
		repo, err := CreateRepo(repository)

		// Validation
		assert.Nil(t, repo, "no repo expected when we get a timeout from github")
		assert.Error(t, err, "an error is expected when get a timeout from github")
		assert.EqualValues(t, "timeout from github", err.Error())
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"Test repository.","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"name":"test-repo","description":"Test repository.","private":true}`,
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "Test repository.",
			Private:     true,
		}

		// Execution
		repo, err := CreateRepo(repository)

		// Validation
		assert.NotNil(t, repo, "a repository was expected")
		assert.NoError(t, err, "no error wa expected")
		assert.EqualValues(t, "test-repo", repo.Name, fmt.Sprintf("name should be %s", repo.Name))
		assert.EqualValues(t, "Test repository.", repo.Description, fmt.Sprintf("description should be %s", repo.Description))
		assert.EqualValues(t, true, repo.Private, fmt.Sprintf("private should be %t", repo.Private))
	})

	t.Run("TestErrorParsingGithubError", func(t *testing.T) {
		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"Test repository.","private":true}`,
			ResponseStatusCode: http.StatusNoContent,
			ResponseBody:       `{"name":"test-repo","description":"Test repository.","private":1}`,
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "Test repository.",
			Private:     true,
		}

		// Execution
		repo, err := CreateRepo(repository)

		// Validation
		assert.Nil(t, repo, "no repo expected when we get a timeout from github")
		assert.Error(t, err, "an error is expected when trying to unmarshal github error")
	})

	t.Run("TestErrorUnmarshalResponse", func(t *testing.T) {
		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"Test repository.","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"name":"test-repo","description":2}`,
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "Test repository.",
			Private:     true,
		}

		// Execution
		repo, err := CreateRepo(repository)

		// Validation
		assert.Nil(t, repo, "no repo expected when we get a timeout from github")
		assert.Error(t, err, "an error is expected when trying to unmarshal github response")
	})
}
