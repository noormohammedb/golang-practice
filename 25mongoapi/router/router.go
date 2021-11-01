package router

import (
	"github.com/gorilla/mux"
	"github.com/noormohammedb/mongoapi/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("get")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("post")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("put")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("delete")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("delete")

	return router
}
