package main

import (
	"fmt"
	"math"
)

func main() {
	// Part 1 -
	// Print out the data types along with the values
	var1 := math.Pi
	var2 := "Hey Gopher!"
	var3 := true
	var4 := 49
	// Print out should be in the following format - e.g var5 := "I am learning Go"
	// Output ->
	// Value: I am learning Go    Data Type: string
	fmt.Printf("Value: %v\tData Type: %T\n", var1, var1)
	fmt.Printf("Value: %v\tData Type: %T\n", var2, var2)
	fmt.Printf("Value: %v\tData Type: %T\n", var3, var3)
	fmt.Printf("Value: %v\tData Type: %T\n", var4, var4)
	// Part 2 -
	var5 := false
	// Guess the output and print out the data types
	if var5 == false {
		// print out after each variable declaration
		var1 = 123.456
		fmt.Printf("Value: %v\tData Type: %T\n", var1, var1)
		// insert print statement here
		var2 = "Yo Gopher"
		fmt.Printf("Value: %v\tData Type: %T\n", var2, var2)
		var3 = true
		fmt.Printf("Value: %v\tData Type: %T\n", var3, var3)
		var4 = 3672 - 982 + 1293
		fmt.Printf("Value: %v\tData Type: %T\n", var4, var4)
		var5 = true
		fmt.Printf("Value: %v\tData Type: %T\n", var5, var5)
	}
}
