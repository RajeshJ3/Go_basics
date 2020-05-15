package main

import "fmt"

func main() {
	// for loop

	// way 1
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// way 2
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz\n")
		} else if i%3 == 0 {
			fmt.Printf("Fizz\n")
		} else if i%5 == 0 {
			fmt.Printf("Buzz\n")
		} else {
			fmt.Println(i)
		}
	}
}
