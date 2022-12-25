package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web Request")
	url := "https://lco.dev/"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Website response : ", response)
	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("content: ", string(databytes))
}
