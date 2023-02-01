package main

import (
	"fmt"
	"time"
)

func main() {
	go GoRoutines("Hello")
	GoRoutines("World")
}

func GoRoutines(s string) {
	count := 6
	for i := 0; i < count; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
