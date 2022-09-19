package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices - makeAppend")
	Scores := make([]int, 2)
	fmt.Printf("The datatype of Scores is %T\n", Scores)

	Scores[0] = 56
	Scores[1] = 78

	Scores = append(Scores, 23, 45, 67)
	fmt.Println(Scores)

	fmt.Println(sort.IntsAreSorted(Scores))

	sort.Ints(Scores)
	fmt.Println(Scores)

	fmt.Println(sort.IntsAreSorted(Scores))

}
