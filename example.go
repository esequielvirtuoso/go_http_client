package main

import (
	"fmt"
	"io/ioutil"

	"github.com/esequielvirtuoso/go_http_client/gohttp"
)

func main() {
	client := gohttp.New()
	response, err := client.Get("https://api.github.com", nil)
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
