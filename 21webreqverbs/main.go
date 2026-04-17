package main

import (
	"fmt"
	"io"
	"strings"

	"net/http"
)

func main() {
	fmt.Println("WebServers")
	PerformGetRequest()
}
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code is :", response.StatusCode)
	fmt.Println("Content Length is :", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	//if err != nil {
	//panic(err)
	//}
	//fmt.Println(string(content))
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is :", byteCount)
	fmt.Println(responseString.String())

}
