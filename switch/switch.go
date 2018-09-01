package main

import "fmt"

func main() {
	var age int
	fmt.Printf("Please enter your age: ")
	fmt.Scan(&age)

	switch {
	case age < 10:
		fmt.Println("You are a child.")
	case age < 19:
		fmt.Println("You are a teen.")
	case age > 19:
		fmt.Println("You are an adult.")
	}

	var username string
	fmt.Printf("Please enter your username: ")
	fmt.Scan(&username)

	switch username {
	// username - octallium
	case "Octallium":
		fmt.Println("Correct Username.")
	case "octallium":
		fmt.Println("Correct Username.")
	default:
		fmt.Println("Wrong Username!")
	}
}
