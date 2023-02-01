package main

import "fmt"

func main() {
	GoRoutines("Hello")
	GoRoutines("World")
}

func GoRoutines(s string) {
	count := 6
	for i := 0; i < count; i++ {
		fmt.Println(s)
	}
}
