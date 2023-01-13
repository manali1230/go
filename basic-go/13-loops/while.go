package main

import "fmt"

func main() {
	fmt.Println("type of while loop in go")
	count := 1
	for count < 10 {
		fmt.Println(count)
		count++
	}
}
