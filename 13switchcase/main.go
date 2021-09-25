package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World!")

	// var day int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Day Number : ")
	readerInput, _ := reader.ReadString('\n')

	// day = strconv.ParseInt(strings.TrimSpace(readerInput))
	day, _ := strconv.ParseInt(strings.TrimSpace(readerInput), 10, 64)

	switch day {
	case 1:
		fmt.Println("sunday")
	case 2:
		fmt.Println("monday")
	case 3:
		fmt.Println("tuesay")
	case 4:
		fmt.Println("wednesday")
	case 5:
		fmt.Println("thursday")
	case 6:
		fmt.Println("friday")
	case 7:
		fmt.Println("saturday")
	default:
		fmt.Println("Error")
	}

}
