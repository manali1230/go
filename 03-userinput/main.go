package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "Welcome to the Rating System"
	fmt.Println(Welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for Service:")

	// comma ok // comma err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the Rating:", input)
	fmt.Printf("The Variable type of input is : %T\n", input)

}
