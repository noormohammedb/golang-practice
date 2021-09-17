package main

import "fmt"

func main() {
	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "mango"

	fmt.Println("fruits, ", fruits)
	fmt.Println("fruits length, ", len(fruits))

	var vegetables = [6]string{"potato", "tomato", "onion"}
	fmt.Println("vegetables, ", vegetables)
	fmt.Println("vegetables length, ", len(vegetables))

}
