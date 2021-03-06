package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()

	webRoute := mux.NewRouter()
	webRoute.NewRoute().HandlerFunc(serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", webRoute))
}

func greeter() {
	fmt.Println("Hello From GO-Lang")

}

func serveHome(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Requested")
	w.Write([]byte("<h1> Hello World From GO-Lang Server </h1>"))

}
