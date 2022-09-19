package main

import "fmt"

func main() {
	fmt.Println("Slices - make")
	Scores := make([]int, 2)
	fmt.Printf("The datatype of Scores is %T\n", Scores)

	Scores[0] = 56
	Scores[1] = 78
	Scores[2] = 14
	fmt.Println(Scores)

}
