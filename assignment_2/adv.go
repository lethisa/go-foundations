package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	symbol := "#"
	var userInput int
	validInput := false
	fmt.Printf("Please enter a number: ")
	fmt.Scan(&userInput)

	for validInput != true {

		if userInput <= 0 {
			fmt.Printf("Enter a valid number greater than 0: ")
			fmt.Scan(&userInput)
		}
		if userInput > 0 {
			validInput = true
		}
	}

	// "%6s\n"
	fmtString := "%" + strconv.Itoa(userInput) + "s\n"
	for i := 1; i <= userInput; i++ {
		fmt.Printf(fmtString, strings.Repeat(symbol, i))
	}
}
