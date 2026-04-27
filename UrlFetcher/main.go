package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching", url, ":", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("\nURL:", url)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Body Length:", len(body))
}

func main() {
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
}
