package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learn Time")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))

	createdTime := time.Date(1999, time.March, 24, 3, 30, 21, 40, time.UTC)
	fmt.Println("Created Time")
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("02-01-2006 Monday 15:04:05"))
}
