package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()

	webRoute := mux.NewRouter()
	webRoute.NewRoute().HandlerFunc(serveHome)

	http.ListenAndServe(":8000", webRoute)
}

func greeter() {
	fmt.Println("Hello From GO-Lang")

}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h2> Hello World From GO-Lang Server </h2>"))

}
