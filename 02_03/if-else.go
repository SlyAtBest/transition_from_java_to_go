package main

import (
	"fmt"
)

func printParity(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}
	fmt.Printf("%v is odd.\n", x)
}

func main() {
	printParity(4)
	printParity(5)
}
