package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post Request")
	PostRequest()

}

func PostRequest() {
	const myurl = "http://localhost:3000/post"
	//json payload [send data]
	requestBody := strings.NewReader(`
	{
		"Name":"Manali",
		"Price":3000,
		"Route":"Train"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
