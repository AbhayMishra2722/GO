package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://go.dev:3000/learn?coursename=reactjs&paymentid=15sa4d4sa54"

func main() {
	fmt.Println("Urls")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qprams := result.Query()
	fmt.Printf("the result of query parameters is %T/n", qprams)
	fmt.Println(qprams["coursename"])

	partsofUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=abhay",
	}
	anotherurl := partsofUrl.String()
	fmt.Println(anotherurl)
}
