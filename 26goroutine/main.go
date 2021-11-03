package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("World!")
}

func greeter(argData string) {
	for i := 0; i < 6; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(argData)
	}
}
