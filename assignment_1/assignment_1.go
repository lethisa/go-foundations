package main

import "fmt"

func main() {
	// ask user for First Name
	var firstName string
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	// ask user for last name
	var lastName string
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	// Your name is: Anil Kulkarni
	fmt.Printf("Your name is: %s %s", firstName, lastName)
}
