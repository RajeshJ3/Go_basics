package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBday method (pointer reciever)
func (p *Person) hasBday() {
	p.age++
}

// getMarried method (pointer reciever)
func (p *Person) getMarried(newLastName string) {
	if p.gender == "f" {
		p.lastName = newLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Rajesh", lastName: "Joshi", city: "Haldwani", gender: "m", age: 21}
	person2 := Person{"Prachi", "Arya", "Haldwani", "f", 19}

	fmt.Println(person1, person2, "\n")

	fmt.Println(person1)
	person1.hasBday()
	fmt.Println(person1, "\n")

	person1.getMarried("Arya")
	person2.getMarried("Joshi")
	fmt.Println(person1, person2)

}
