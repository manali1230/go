package router

import (
	"github.com/gorilla/mux"
	"github.com/manali1230/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/get/movies", controller.ResponseAllMovies).Methods("GET")
	router.HandleFunc("/post/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/put/movies/{id}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/deleteone/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/deleteall/movies", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
