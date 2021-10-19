package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev/learn?course=golang&paymentid=blabla&price=0"

func main() {
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme, "", result.Host, result.Path, result.RawQuery)

	qParams := result.Query()

	fmt.Printf("\ntype of qParams : %T\n", qParams)

	fmt.Println(qParams["paymentid"])

	for key, val := range qParams {
		fmt.Print(key, " : ", val, ", ")
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/lorem",
		RawQuery: "search=ipsum",
	}

	fmt.Println("\n", partsOfUrl)
	fmt.Println(partsOfUrl.String())

}
