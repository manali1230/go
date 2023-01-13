package main

import "fmt"

func main() {
	fmt.Println("If-ELSE")

	num := 20
	var result string

	if num < 20 {
		result = "number is less than 20"
	} else if num > 20 {
		result = "number is greater than 20"
	} else {
		result = "number is equal to 20"
	}

	fmt.Println(result)
}
