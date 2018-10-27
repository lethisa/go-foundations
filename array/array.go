package main

import "fmt"

func main() {
	friends := [4]string{"John", "Jane", "Joe", "Janet"}
	fmt.Println("Length:", len(friends))
	fmt.Println("Capacity:", cap(friends))
	for i := 0; i < len(friends); i++ {
		fmt.Printf("Index: %v\tValue: %v\n", i, friends[i])
	}
}
