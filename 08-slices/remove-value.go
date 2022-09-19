package main

import "fmt"

func main() {
	fmt.Println("Remove a value from the Slices based on index")
	var langList = []string{"Hindi", "English", "Marathi", "Bengali", "Gujarati"}
	fmt.Println(langList)
	var index int = 2
	langList = append(langList[:index], langList[index+1:]...)
	fmt.Println(langList)
}
