package main

import "fmt"

func main() {
	fmt.Println("GOTO")
	num := 1
	for num < 10 {
		if num == 5 {
			num++
			goto num
		}
		fmt.Println(num)
		num++
	}

num:
	{
		fmt.Println("This is GOTO!")
	}
}
