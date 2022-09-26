package main

import "fmt"

func main() {
	fmt.Println("Loops in GO")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	fmt.Println("Way 1")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// j is index
	fmt.Println("Way 2")
	for j := range days {
		fmt.Println(days[j])
	}

	fmt.Println("Way 3")
	for index, value := range days {
		fmt.Printf("Index %v : value %v\n", index, value)
	}

}
