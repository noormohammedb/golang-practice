package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("web requests\nGET")
	GetMethod("http://localhost:8000/get")
	fmt.Println("POST")
	PostJsonMethod("http://localhost:8000/post")
}

func GetMethod(url string) {

	httpResponse, err := http.Get(url)

	checkNotNill(err)

	defer httpResponse.Body.Close()

	fmt.Println("response status : ", httpResponse.Status)
	fmt.Println("content type : ", httpResponse.Header["Content-Type"])
	fmt.Println("content length : ", httpResponse.ContentLength)

	responseContent, _ := ioutil.ReadAll(httpResponse.Body)

	// fmt.Println("response body : ", string(responseContent))

	var responseString strings.Builder
	responseString.Write(responseContent)

	fmt.Println("response body : ", responseString.String())

	// fmt.Printf("type of responseSTring : %T\n", responseString)
}

func PostJsonMethod(url string) {

	reqBody := strings.NewReader(`{"key1":"value1"}`)

	fmt.Printf("\nType of reqBody : %T\n", reqBody)

	/*
		post body is type Reader, Reader type available in both io.Reader and strings.Reader
	*/
	response, err := http.Post(url, "application/json", reqBody)
	checkNotNill(err)

	defer response.Body.Close()

	contest, _ := ioutil.ReadAll(response.Body)

	fmt.Println("response : ", string(contest))
}

func checkNotNill(err error) {
	if err != nil {
		panic(err)
	}
}
