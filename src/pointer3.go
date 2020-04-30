package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("&x == &x, &x == &y, &x == nil", &x == &x, &x == &y, &x == nil)

	// Function local variable persistence
	var p = f()
	fmt.Println("Address of function local variable", p)
	fmt.Println("Value of function local variable", *p)

	fmt.Println("Evaluate f() == f()", f() == f())
}

func f() *int {
	v := 1

	return &v
}
