package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CourseName string
	Tutor      []string `json:"tutor"`
}

func main() {
	fmt.Println("Decode Json")
	DecodeJson()
}

func DecodeJson() {
	JsonDataFromWeb := []byte(`
	{
		"name": "Go Course",
		"price": 3000,
		"CourseName": "Go",
		"tutor": [
				"Manoj,Saroj"
		]
	}
	`)
	var AllCourses course
	CheckValid := json.Valid(JsonDataFromWeb)

	if CheckValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(JsonDataFromWeb, &AllCourses)
		fmt.Printf("\n%#v\n", AllCourses)
	} else {
		fmt.Println("Json Not Valid")
	}

	// add data to key value
	var online map[string]interface{}
	json.Unmarshal(JsonDataFromWeb, &online)
	fmt.Printf("\n%#v\n", online)
	for key, value := range online {
		fmt.Printf("Key : %v and value : %v => datatype : %T\n", key, value, value)
	}
}
