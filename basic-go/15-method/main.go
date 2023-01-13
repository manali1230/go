package main

import "fmt"

func main() {
	fmt.Println("Methods in Go Lang")
	// there is no inheritance in go language ;no super or parent

	manali := Users{"Manali Jain", "manali@gmail.com", false, 23}
	fmt.Println(manali)
	fmt.Printf("Manali's details are : %+v\n", manali)
	fmt.Printf("User Name is %v and email is %v\n", manali.Name, manali.Email)

	manali.GetStatus()
	manali.SendMail()

	fmt.Printf("User Name is %v and email is %v\n", manali.Name, manali.Email)
}

type Users struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u Users) GetStatus() {
	fmt.Println("Status of user is:", u.Status)
}

func (u Users) SendMail() {
	u.Email = "test@email.com"
	fmt.Println("Email of user is:", u.Email)
}
