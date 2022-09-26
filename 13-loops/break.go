package main

import "fmt"

func main() {
	fmt.Println("Break")
	num := 1
	for num < 10 {
		if num == 5 {
			num++
			break
		}
		fmt.Println(num)
		num++
	}
}
