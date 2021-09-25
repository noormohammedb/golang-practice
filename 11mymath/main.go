package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main() {
	// seed := time.Now().UnixNano()
	// fmt.Println("seed", seed)
	// rand.Seed(seed)
	// fmt.Println(rand.Intn(int(seed)))

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(2))
	fmt.Println(myRandomNumber)
}
