package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test signals"}
var wg sync.WaitGroup // pointer
var mutex sync.Mutex  // pointer

func main() {
	websiteList := []string{
		"https://google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://go.dev",
	}

	for _, web := range websiteList {
		fmt.Printf("%s Endpoint !\n", web)
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("%s Endpoint failed !\n", endpoint)
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		fmt.Printf("%d is the status code for %s\n", response.StatusCode, endpoint)
	}
}
