package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and Case in go")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Your Dice Number is : ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("unloack and move with 1 step")
	case 2:
		fmt.Println("move with 2 steps")
	case 3:
		fmt.Println("move with 3 steps")
		fallthrough // by default each case will not run so fallthrough is used
	case 4:
		fmt.Println("move with 4 steps")
		fallthrough
	case 5:
		fmt.Println("move with 5 steps")
		fallthrough
	case 6:
		fmt.Println("move with 6 steps and roll the dice again")
	default:
		fmt.Println("You are out")
	}

}
