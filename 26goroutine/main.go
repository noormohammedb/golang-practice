package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	// greeter("World!")
	websiteList := []string{
		"https://go.dev/",
		"https://lco.dev/",
		"https://google.com/",
		"https://github.com/",
		"https://gmail.com/",
	}
	for _, webUri := range websiteList {
		// getStatusCode(webUri)
		go getStatusCode(webUri)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(argData string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(300 * time.Millisecond)
// 		fmt.Println(argData)
// 	}
// }

func getStatusCode(uriData string) {
	defer wg.Done()

	res, err := http.Get(uriData)

	if err != nil {
		fmt.Println("Error in api call ", uriData)
	}

	fmt.Printf("%d %s\n", res.StatusCode, uriData)

}
