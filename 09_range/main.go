package main

import "fmt"

func main() {
	ids := []int{32, 43, 65, 83, 24, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using id
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for _, v := range emails {
		fmt.Println("E-mails: ", v)
	}

}
