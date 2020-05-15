package main

import "fmt"

func main() {
	// Define map

	// way 1
	emails0 := make(map[string]string)

	// Assign key values
	emails0["Bob"] = "bob@gmail.com"
	emails0["Sharon"] = "sharon@gmail.com"
	emails0["Mike"] = "mike@gmail.com"

	// way 2
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// delete
	delete(emails, "Bob")
	fmt.Println(emails)
}
