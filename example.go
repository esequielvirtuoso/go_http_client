package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/esequielvirtuoso/go_http_client/gohttp"
)

var (
	gitHubHttpClient = getGitHubClient()
)

func getGitHubClient() gohttp.HttpClient {
	client := gohttp.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getRequest()

}

func getRequest() {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")

	response, err := gitHubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {
	response, err := gitHubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
