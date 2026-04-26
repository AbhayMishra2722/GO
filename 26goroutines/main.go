package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // usually is a pointer
var mut sync.Mutex    // pointer

func main() {
	//greeter("Hello")
	//greeter("World")
	websiteList := []string{
		"https://github.com/AbhayMishra2722",
		"https://leetcode.com/u/DcYnltYXUU/",
		"https://www.linkedin.com/in/abhay-mishra-93404424a/",
		"https://www.codechef.com/users/abhaymishra4",
		"https://www.credly.com/users/abhay-mishra.9f580d19/edit#credly",
	}
	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCode(web)

	}
	wg.Wait()
	fmt.Println(signals)

}

// func greeter(s string) {
//
//	/	for i := 0; i < 6; i++ {
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("200 %d status code for website %s \n", result.StatusCode, endpoint)
	}

}
