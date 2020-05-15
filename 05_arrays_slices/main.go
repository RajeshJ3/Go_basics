package main

import "fmt"

func main() {
	// Arrays
	var fruitArr1 [2]string

	// Assign values
	fruitArr1[0] = "Apple"
	fruitArr1[1] = "Orange"

	// Declare and assign
	fruitArr2 := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr2)
	fmt.Println(fruitArr2[1])

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grapes"}

	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[len(fruitSlice)-1])
	fmt.Println(fruitSlice[:len(fruitSlice)])
}
