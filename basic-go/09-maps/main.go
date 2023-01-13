package main

import "fmt"

func main() {
	fmt.Println("Maps")

	ComplangList := make(map[string]string)

	ComplangList["PY"] = "Python"
	ComplangList["JS"] = "Javascript"
	ComplangList["RB"] = "Ruby"

	fmt.Println("List of Compute Languages : ", ComplangList)
	fmt.Println("PY stands for : ", ComplangList["PY"])

	// delete value from map
	delete(ComplangList, "RB")
	fmt.Println("List of Compute Languages : ", ComplangList)

	// Loops
	for key, value := range ComplangList {
		fmt.Printf("For key %v the value is %v\n", key, value)
	}
}
