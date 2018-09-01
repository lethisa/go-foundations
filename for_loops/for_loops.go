package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
		// Pass 1 - 0
		// Pass 2 - 1 ...
		// Pass n - 10 ... i = 11
	}
	fmt.Println()

	// j := 11
	// for ; j <= 20; j++ {
	// 	fmt.Printf("%d ", j)
	// }
	// fmt.Println()

	j := 11
	for {
		fmt.Printf("%d ", j)
		j++
		if j > 20 {
			break
		}
	}
	fmt.Println()

	for k := 30; k >= 20; k-- {
		fmt.Printf("%d ", k)
	}
	fmt.Println()

	for l := 1; l <= 10; l++ {
		if l == 5 {
			continue
		}
		fmt.Printf("%d ", l)
	}
	fmt.Println()
}
