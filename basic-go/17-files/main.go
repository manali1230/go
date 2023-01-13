package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in GO")
	content := "This is my File !"
	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is", length)
	defer file.Close()
}
