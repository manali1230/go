package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Welcome := "Welcome to the Rating !!"
	fmt.Println(Welcome)
	fmt.Println("Rate us between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating", input)
	fmt.Printf("variable type of input is %T\n", input) // rating in string

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("your Rating is ", numRating)
		fmt.Printf("Variable Type %T\n", numRating) // rating converted to float
	}
}
