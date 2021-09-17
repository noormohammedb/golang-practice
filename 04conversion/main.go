package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter Rating : ")

	reader := bufio.NewReader(os.Stdin)
	readerData, _ := reader.ReadString('\n')

	fmt.Println("Thanks For Rating,", readerData)

	dataParsed, parseError := strconv.ParseFloat(strings.TrimSpace(readerData), 64)
	if parseError != nil {
		fmt.Println("caught error in parsing data")
		fmt.Println(parseError)

	} else {
		newValue := dataParsed + 4
		fmt.Println("added 4 to your rating", newValue)
	}
}
