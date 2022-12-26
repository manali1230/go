package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://103.133.123.157:3000/dev?language=go&id=abc123"

func main() {
	fmt.Println("URL")

	fmt.Println("My URL is : ", myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println("Host : ", result.Host)
	fmt.Println("Protocol :", result.Scheme)
	fmt.Println("Port : ", result.Port())
	fmt.Println("Path : ", result.Path)
	fmt.Println("Query : ", result.RawQuery)

	queryParameter := result.Query()
	fmt.Println("Query Parameter : ", queryParameter)

	fmt.Println("ID : ", queryParameter["id"])
	fmt.Println("Language : ", queryParameter["language"])

	for _, val := range queryParameter {
		fmt.Println("Parameter : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "103.133.123.157",
		Path:     "/prod",
		RawQuery: "user=Manali",
	}

	getUrl := partsOfUrl.String()
	fmt.Println(getUrl)
}
