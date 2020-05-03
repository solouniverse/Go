package main

import "fmt"

func main() {
	var x int
	fmt.Print("Enter x value: ")
	fmt.Scan(&x)

	if x % 2 == 0 {
		fmt.Printf("%d is even\n", x)
	} else {
		fmt.Printf("%d is odd\n", x)
	}
}
