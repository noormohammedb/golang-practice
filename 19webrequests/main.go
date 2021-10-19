package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("Web Request")
	httpResponse, err := http.Get(url)
	checkNotNill(err)

	fmt.Printf("response type : %T\n", httpResponse)
	fmt.Printf("body type : %T\n", httpResponse.Body)

	// fmt.Println(httpResponse.Body)

	defer httpResponse.Body.Close()

	httpResponseContent, err := ioutil.ReadAll(httpResponse.Body)
	checkNotNill(err)

	fmt.Println(string(httpResponseContent))
}

func checkNotNill(err error) {
	if err != nil {
		panic(err)
	}
}
