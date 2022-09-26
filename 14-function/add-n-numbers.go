package main

import "fmt"

func main() {
	result := n_addition(1, 2, 3, 4, 7)
	fmt.Println(result)
}

func n_addition(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
