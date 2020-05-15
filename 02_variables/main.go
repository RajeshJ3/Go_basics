package main

import "fmt"

func main() {
	var x int
	x = 10

	var y int = 20

	fmt.Println(x, y)

	name, age := "Rajesh", 21
	fmt.Println(name, age)
	fmt.Printf("%T\n", name)
}
