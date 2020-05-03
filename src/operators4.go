package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter x value:")
	fmt.Scan(&x)

	if x % 2 == 0 {
		fmt.Printf("%d is even\n", x)
		if x % 5 == 0 {
			fmt.Printf("%d can be divisible by 5\n", x)
		}
	} else {
		fmt.Printf("%d is odd\n", x)
	}
}
