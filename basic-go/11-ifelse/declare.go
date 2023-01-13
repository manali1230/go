package main

import "fmt"

func main() {
	fmt.Println("Declare variable with if-else")

	if num := 9; num < 5 {
		fmt.Println("number is less than 5")
	} else {
		fmt.Println("number is greater than or equal to 5")
	}
}
