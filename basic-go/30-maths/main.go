package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {

	// Math
	var num1 int = 2
	var num2 float64 = 3.4
	fmt.Println("Sum of numbers : ", num1+int(num2))

	// Random Number
	// rand.Seed(time.Now().UnixMicro())
	// fmt.Println(rand.Intn(5) + 1)

	// Crypto
	randNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randNum)
}
