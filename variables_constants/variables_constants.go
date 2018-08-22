package main

import "fmt"

func main() {
	// var - keyword
	var username string
	username = "Octallium"
	fmt.Println("Username:", username)

	// short-hand or the inference type method of declaration
	firstName := "Anil"
	lastName := "Kulkarni"
	fmt.Println("Full Name:", firstName, lastName)

	// declaring multi-line variables
	var (
		a string = "tic"
		b int    = 4
		c bool   = true
	)
	fmt.Println(a, b, c)

	// declaring multiple variables on the same line
	var d, e, f = "I am a string", 9.0, true
	fmt.Printf("d: %v, e: %.2f, f: %v\n", d, e, f)

	// changing the value of the variable 'd'
	d = "is now a new string"
	fmt.Println("d:", d)

	// declaring a constant
	const API_KEY = "dfbkj.adsb.132423"
	fmt.Println(API_KEY)
}
