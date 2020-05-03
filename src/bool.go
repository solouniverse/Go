package main

import "fmt"

func main() {
	var x bool
	fmt.Println("x default bool value:", x)

	x = 1 < 5
	fmt.Println("1 < 5", x)

	x = 1 > 5
	fmt.Println("1 > 5", x)
}
