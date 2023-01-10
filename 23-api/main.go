package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// model
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper or a funtion which encrypt data only when ID is given
func (c *Course) isEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {

}

// backend route
func backend(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
