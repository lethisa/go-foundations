package main

import (
	"fmt"
	"strings"
)

func main() {

	var num int
	fmt.Printf("Please enter a number: ")
	fmt.Scan(&num)
	symbol := "#"
	validInput := false

	for validInput != true {
		if num <= 0 {
			fmt.Printf("Please enter a valid number greater than 0: ")
			fmt.Scan(&num)
		}
		if num > 0 {
			validInput = true
		}
	}

	for i := 1; i <= num; i++ {
		fmt.Printf("%v\n", strings.Repeat(symbol, i))
	}
}
