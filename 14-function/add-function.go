package main

import "fmt"

func main() {
	result := Addition(2, 3)
	fmt.Println(result)
}

func Addition(num1 int, num2 int) int {
	return num1 + num2
}
