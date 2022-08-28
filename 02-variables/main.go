package main

import "fmt"

// numberOfSingup := 12000 not allowed
// var numberOfSingup = 12000 allowed

const LoginToken = "abcdefgh" // this is a public variable that's why LoginToken "L" is capital

func main() {
	var username string = "Manali"
	fmt.Println(username)
	fmt.Printf("The variable type is : %T \n", username)

	var isSignUp bool = true
	fmt.Println(isSignUp)
	fmt.Printf("The variable type is : %T \n", isSignUp)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("The variable type is : %T \n", smallval)

	var intval int = -1
	fmt.Println(intval)
	fmt.Printf("The variable type is : %T \n", intval)

	var smallFloatVal float32 = 255.12344
	fmt.Println(smallFloatVal)
	fmt.Printf("The variable type is : %T \n", smallFloatVal)

	var FloatVal float64 = 255.12345123451234
	fmt.Println(FloatVal)
	fmt.Printf("The variable type is : %T \n", FloatVal)

	// default values and some aliases

	var count int
	fmt.Println(count)
	fmt.Printf("The variable type is : %T \n", count)

	var name string
	fmt.Println(name)
	fmt.Printf("The variable type is : %T \n", name)

	// implicit type

	var website = "test.web.net"
	fmt.Println(website)
	fmt.Printf("The variable type is : %T \n", website)

	// no var style [:= is known as walrus operator]

	numberOfSingup := 12000
	fmt.Println(numberOfSingup)
	fmt.Printf("The variable type is : %T \n", numberOfSingup)

	area := 12000.1234
	fmt.Println(area)
	fmt.Printf("The variable type is : %T \n", area)

	// public variable\

	fmt.Println(LoginToken)
	fmt.Printf("The variable type is : %T \n", LoginToken)
}
