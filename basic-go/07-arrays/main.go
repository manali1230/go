package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Arrays")

	var langList [3]string

	langList[0] = "Hindi"
	langList[2] = "English"

	fmt.Println("Value of language list is", langList)
	fmt.Println("Length of language list is", len(langList))

	var compLangList = [2]string{"C","C++"}
	
	fmt.Println("Value of Computer language list is", compLangList)
	fmt.Println("Length of Computer language list is", len(compLangList))
}