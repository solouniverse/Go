package main

import "fmt"

func main() {
	// tuple assignments
	x, y := 10, 20
	fmt.Printf("x = %d, y = %d\n", x, y)

	// swapping values of x, y
	x, y = y, x
	fmt.Println("x =", x, "y =", y)
}
