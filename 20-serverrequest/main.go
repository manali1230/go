package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Server Request")
	GetRequest()

}

func GetRequest() {
	const myurl = "http://localhost:3000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	bytecount, _ := responseString.Write(content)

	fmt.Println("ByteCount : ", bytecount)
	fmt.Println(responseString.String())

	// fmt.Println(content)     // in bytes
	// fmt.Println(string(content))
}
