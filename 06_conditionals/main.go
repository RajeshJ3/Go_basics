package main

import "fmt"

func main() {

	// if else
	x, y := 10, 10
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if clause
	color := "black"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "black" {
		fmt.Println("Color is black")
	} else {
		fmt.Println("Unknown color!")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "black":
		fmt.Println("Color is black")
	default:
		fmt.Println("Unknown color!")
	}
}
