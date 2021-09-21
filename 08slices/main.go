package main

import "fmt"

func main() {
	// var fruitList = []string{"mango", "apple", "banana"}
	fruitList := []string{"mango", "apple", "banana"}

	fmt.Printf("type of fruitList, %T \n", fruitList)

	fmt.Println(fruitList)

	// fruitList = append(fruitList, []string{"orange", "avocado"}...)
	fruitList = append(fruitList, "orange", "avocado")

	fmt.Println("full list: ", fruitList)

	fruitList = append(fruitList[:2], fruitList[3:]...) // remove item 2

	fmt.Println(fruitList)
	// fmt.Println(fruitList[2:5])
}
