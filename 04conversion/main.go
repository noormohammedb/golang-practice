package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter Rating : ")

	reader := bufio.NewReader(os.Stdin)
	readerData, _ := reader.ReadString('\n')

	fmt.Println("Thanks For Rating, ", readerData)

	dataParsed, parseError := strconv.ParseFloat(readerData, 64) // trim spaces before parsing
	if parseError != nil {
		fmt.Println("caught error")
		fmt.Println(parseError)

	} else {
		newValue := dataParsed + 1
		fmt.Println(newValue)
		fmt.Println(dataParsed)
	}
}
