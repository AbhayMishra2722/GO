package main

import (
	"fmt"
	"io"
	"strings"

	"net/http"
	"net/url"
)

func main() {
	fmt.Println("WebServers")
	//	PerformGetRequest()
	// PerformPostJsonRequest()
	performPostformRequest()
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
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake JSON payload
	requestBody := strings.NewReader(`
	{
        "coursename":"Lets go with Golang",
        "price":0,
		"platform":"youtube"

	}
`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}
func performPostformRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "abhay")
	data.Add("lastname", "mishra")
	data.Add("course", "golang")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
