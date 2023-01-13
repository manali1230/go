package main

import "fmt"

func main() {
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	count := 5
	for i := 0; i < count; i++ {
		defer fmt.Println(i)
	}
}
