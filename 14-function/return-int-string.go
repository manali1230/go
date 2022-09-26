package main

import "fmt"

func main() {
	result, message := n_addition(1, 2, 3, 4, 7)
	fmt.Println(result)
	fmt.Println(message)
}

func n_addition(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This is addition function!"
}
