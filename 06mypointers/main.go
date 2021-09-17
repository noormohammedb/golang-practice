package main

import "fmt"

func main() {
	var myPtr *int
	fmt.Println("myPtr value: ", myPtr)

	var myRealNum = 85
	// var myNewPtr = &myRealNum
	myNewPtr := &myRealNum

	fmt.Println("myRealNum value: ", myRealNum)
	fmt.Println("myRealNum address: ", &myRealNum)
	fmt.Println("myNewPtr value: ", myNewPtr)
	fmt.Println("myNewPtr address: ", &myNewPtr)
	fmt.Println("myNewPtr reference: ", *myNewPtr)

	*myNewPtr = *myNewPtr + 44
	fmt.Println("after operation")

	fmt.Println("myNewPtr reference: ", *myNewPtr)
	fmt.Println("myRealNum value: ", myRealNum)
}
