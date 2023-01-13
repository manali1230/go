package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Pointers Usage")
	
	var num = 3 
    var ptr = &num  // & means reference

	fmt.Println("pointer value is", ptr)  // pointer is storing reference of num variable
	fmt.Println("value of pointer is", *ptr)  // by using *ptr we can print the actual value of refenced variable

	*ptr = *ptr * 10
	fmt.Println("value of num is", num)   // any value change in pointer is reflected in the referenced variable also
}