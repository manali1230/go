package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in GO")
	content := "This is my File !"
	file, err := os.Create("./file.txt")
	checkNilerr(err)
	length, err := io.WriteString(file, content)
	checkNilerr(err)
	fmt.Println("length is", length)
	defer file.Close()
	ReadFile("./file.txt")
}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilerr(err)
	// fmt.Println("Data of file ", databyte)
	fmt.Println("Data inside file - ", string(databyte))
}
func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
