package main

import (
	"fmt"
	"sort"
)

func main() {
	// var fruitList = []string{"mango", "apple", "banana"}
	fruitList := []string{"mango", "apple", "banana"}

	fmt.Printf("type of fruitList, %T \n", fruitList)

	fmt.Println(fruitList)

	// fruitList = append(fruitList, []string{"orange", "avocado"}...)
	fruitList = append(fruitList, "orange", "avocado")

	fmt.Println("full list: ", fruitList)

	var index int = 2

	fruitList = append(fruitList[:index], fruitList[index+1:]...) // remove item 2

	fmt.Println(fruitList)
	// fmt.Println(fruitList[2:5])

	scores := make([]int, 3)

	scores[0] = 11
	scores[1] = 44
	scores[2] = 34
	// scores[3] = 22 // getting error, index out of range

	fmt.Println("dynamic slice")
	fmt.Println(scores)
	// fmt.Println(len(scores))

	scores = append(scores, 77, 33, 00, 22) // allocate more spaces to scores
	// fmt.Println(append(scores, 77, 33, 00, 22))
	fmt.Println(scores)

	// fmt.Println(sort.IntsAreSorted(scores))
	fmt.Println("sorting")
	sort.Ints(scores)
	fmt.Println(scores)

	// fmt.Println(sort.IntsAreSorted(scores))
}
