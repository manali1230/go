package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go Modules Backend")
	r := mux.NewRouter()
	r.HandleFunc("/", Backend).Methods("Get")
	log.Fatal(http.ListenAndServe(":1000", r))
}

func Backend(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go Backend</h1>"))
}
