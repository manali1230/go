package main

import "fmt"

func main() {
	fmt.Println("Check Even/Odd")
	num := 10
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
