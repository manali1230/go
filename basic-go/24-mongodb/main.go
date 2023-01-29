package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/manali1230/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API")
	fmt.Println("Server is started ...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":1000", r))
	fmt.Println("Listening at port 1000 ...")

}
