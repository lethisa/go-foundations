package main

import "fmt"

func main() {
	// if - else
	num := 11
	if num <= 10 {
		fmt.Println("Less than or equal to 10.")
	} else if num > 10 {
		fmt.Println("Greater than 10.")
	}

	username := "Octallium"
	username = "octallium"
	if username == "Octallium" || username == "octallium" {
		fmt.Println("Username is correct.")
	} else {
		fmt.Println("Wrong username.")
	}

	age := 14
	age = 15
	if username == "octallium" && age == 14 {
		fmt.Println("Username is correct.")
	} else {
		fmt.Println("Wrong username.")
	}
}
