package main

import (
	"fmt"
)

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Staurday"}

	fmt.Println(days)

	// normal forLoop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// range forLoop
	for i := range days {
		fmt.Println(days[i])
	}

	// like forEach loop
	for index, day := range days {
		fmt.Printf("%v. %v\n", index, day)
	}

	dumpValue := 1

	// like whileLook
	for dumpValue <= 10 {
		if dumpValue == 5 {
			dumpValue++
			continue
			// break
		}
		fmt.Print(dumpValue, ", ")
		dumpValue++
	}
}
