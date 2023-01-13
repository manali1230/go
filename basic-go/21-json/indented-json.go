package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name       string
	Price      int
	CourseName string
	Tutor      []string
}

func main() {
	fmt.Println("Json Data")
	EncodeJson()
}

func EncodeJson() {
	Courses := []course{
		{"Python Course", 2000, "Python", []string{"Rina", "Mina"}},
		{"Go Course", 3000, "Go", []string{"Manoj,Saroj"}},
	}

	// package above data into json data
	jsondata, err := json.MarshalIndent(Courses, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Println(jsondata)   in bytes
	fmt.Printf("%s\n", jsondata)
}
