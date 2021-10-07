/*
	calc function is passed to sayHai anonymus function as argument

	reference
	https://golangdocs.com/anonymous-functions-in-golang/
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	sayHello()

	sayHai := func(callBack func(numOne, numTwo int) (int, string)) { // anonymous
		numbers := []int{4, 5}
		fmt.Print("adding ")
		fmt.Println(numbers)
		fmt.Println(callBack(numbers[0], numbers[1]))
	}

	sayHai(calc) // normal function declaration not allowed in main

	// values := []int{4, 6, 2, 4}
	// sum := proCalc(values...)
	sum := proCalc(5, 6, 7)

	fmt.Println(sum)
}

func calc(numOne, numTwo int) (int, string) {
	return numOne + numTwo, "Added"
}

func proCalc(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sayHello() {
	fmt.Println("Hello")
}
