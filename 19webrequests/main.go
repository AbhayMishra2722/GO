package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev"

func main() {
	fmt.Println("Web Request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type : %T /n", response)

	defer response.Body.Close() // callers responsibility to close the connectiion

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(data)
	fmt.Println(content)
}
