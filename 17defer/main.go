package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Print("Hello ")
	myDefer()
}

func myDefer() {
	for i := 1; i < 4; i++ {
		defer fmt.Print(i, " ")
	}
}
