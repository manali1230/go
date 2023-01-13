package main

import "fmt"

func main() {
	fmt.Println("Continue")
	num := 1
	for num < 10 {
		if num == 5 {
			num++
			continue
		}
		fmt.Println(num)
		num++
	}
}
