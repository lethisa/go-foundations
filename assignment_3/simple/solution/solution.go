package main

import "fmt"

func main() {

	var num int
	fmt.Printf("Please enter a number: ")
	fmt.Scan(&num)
	//   #
	//  ##
	// ###
	symbol := "#"

	if num < 0 {
		fmt.Printf("Please enter a valid number greater than 0: ")
		fmt.Scan(&num)
	}

	for i := 1; i <= num; i++ {
		// Pass 1 -> i = 1
		// Pass 2 -> i = 2
		// Pass 3 -> i = 3
		// Pass 4 -> i = 4
		for j := 1; j <= i; j++ {
			// Pass 1 -> j = 1
			// Output -> #
			// Pass 2 -> j = 2
			// Pass 3 -> j = 1
			// Output -> ##
			// Pass 4 -> j = 2
			// Pass 5 -> j = 3
			// Pass 6 -> j = 1
			// Output -> ###
			// Pass 7 -> j = 2
			// Pass 8 -> j = 3
			// Pass 9 -> j = 4
			fmt.Printf("%v", symbol)
		}
		fmt.Println()
	}
}
