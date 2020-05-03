package main

import "fmt"

func main() {
	x, y, z := togo()
	fmt.Println("x =", x, "y =", y, "z =", z)

	p, _, r := togo()
	fmt.Println("p =", p, "r =", r)
}

// multiple return values are parenthesized
func togo() (int, int, int) {
	return 1, 2, 3
}
