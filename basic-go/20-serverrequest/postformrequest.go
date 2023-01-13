package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Post Form Request")
	PostFormRequest()

}

func PostFormRequest() {
	const myurl = "http://localhost:3000/post"
	//form data
	data := url.Values{}

	data.Add("Name", "Manali Jain")
	data.Add("City", "Gurgaon")
	data.Add("email", "manali@email.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
