package main

import "fmt"

func main() {

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "10 user logins"
	}
	fmt.Println(result)

	if 8%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 20; num < 50 {
		fmt.Println("number in smaller")
	} else {
		fmt.Println("number in bigger")
	}

}
