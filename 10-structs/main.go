package main

import "fmt"

func main() {
	fmt.Println("Structs in Go Lang")
	// there is no inheritance in go language ;no super or parent

	manali := Users{"Manali Jain", "manali@gmail.com", true, 23}
	fmt.Println(manali)
	fmt.Printf("Manali's details are : %+v\n", manali)
	fmt.Printf("User Name is %v and email is %v\n", manali.Name, manali.Email)
}

type Users struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
