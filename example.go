package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/esequielvirtuoso/go_http_client/gohttp"
)

var (
	gitHubHttpClient = getGitHubClient()
)

func getGitHubClient() gohttp.Client {
	builder := gohttp.NewBuilder()
	builder.SetConnectionTimeout(2 * time.Second)
	builder.SetRequestTimeout(50 * time.Millisecond)
	// builder.DisableTimeouts(true)

	// commonHeaders := make(http.Header)
	// commonHeaders.Set("Authorization", "Bearer ABC-123")

	// builder.SetHeaders(commonHeaders)

	return builder.Build()
}

func main() {
	getRequest()

}

func getRequest() {
	headers := make(http.Header)
	// headers.Set("Authorization", "Bearer ABC-123")

	response, err := gitHubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	// Using custom response
	var user User
	if err := response.UnmarshalJson(&user); err != nil {
		panic(err)
	}
	fmt.Print(user.FirstName)

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

	fmt.Println(response.StatusCode())

	fmt.Println(string(response.Bytes()))
}
