package main

import (
	"fmt"
)

// viewCount := 43 // valarus operator not allowed outside method

const Token = "lorem ipsum dollor" // capitel first letter make this a public entity

func main() {
	var name string = "blabla"
	fmt.Println(name)
	fmt.Printf("type of variable name : %T \n", name)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("type of variable IsLoggedin : %T \n", isLoggedin)

	var numint uint8 = 255
	fmt.Println(numint)
	fmt.Printf("type of variable numint : %T \n", numint)

	var numfloat float64 = 255.12345678901234
	fmt.Println(numfloat)
	fmt.Printf("type of variable numfloat : %T \n", numfloat)

	// default values and some aliases
	var uninitialised int
	fmt.Println(uninitialised)
	fmt.Printf("type of variable uninitialised : %T \n", uninitialised)

	// implicit type
	var lorem = "ipsum"
	fmt.Println(lorem)

	// no var style
	userCount := 333.0 // valarus operator
	fmt.Println(userCount)

	fmt.Println(Token)
	fmt.Printf("type of variable Token : %T \n", Token)

}
