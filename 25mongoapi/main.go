package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/noormohammedb/mongoapi/router"
)

func main() {
	fmt.Println("golang mongodb crud api")
	fmt.Println("Server is getting started...")
	rtr := router.Router()
	log.Fatal(http.ListenAndServe(":8080", rtr))
	fmt.Println("Listening at port 8080 ...")

}
