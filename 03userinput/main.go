package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter The Rating : ")
	readerInput, _ := reader.ReadString('\n')

	fmt.Println("Your rating is : ", readerInput)
	fmt.Printf("Type Of ReaderInput : %T\n", readerInput)
}
