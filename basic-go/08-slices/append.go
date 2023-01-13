package main

import "fmt"

func main() {
	fmt.Println("Append Slices")
	var langList = []string{"Hindi", "English", "Marathi"}

	langList = append(langList, "Bengali", "Gujarati")
	fmt.Println(langList)

	langList = append(langList[1:]) // 0th position is excluded
	fmt.Println(langList)

	langList = append(langList[1:4]) // 0th & 4th position are not included also last position is always non-inclusive
	fmt.Println(langList)

	langList = append(langList[:1]) // as i am always adding langList variable it's value is change means it's 0th position is now Marathi not Hindi
	fmt.Println(langList)
}
