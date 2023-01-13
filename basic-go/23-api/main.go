package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
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

func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Courses")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	param := mux.Vars(r)
	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseID == param["id"] {
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with requested id")
	return
}

func CreateOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Courses")
	w.Header().Set("Content-Type", "application/json")
	// body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add some data")
	}

	// body is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("add data in json")
		return
	}

	// generate unique course id and convert it into string then add course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Courses")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	param := mux.Vars(r)

	// update id
	for index, course := range courses {
		if course.CourseID == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = param["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}
	}
}
