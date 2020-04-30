package main

import "fmt"

func main() {
	var pointer *int
	fmt.Println("Pointer address: ", pointer)
	callByRef(&pointer)
	fmt.Println("Pointer address: ", pointer)
	fmt.Println("Pointer value: ", *pointer)
}

func callByRef(ptr **int) {
	// x local variable
	// life and scope of x must be within this block
	x := 10

	// when assigned x to pointer the local variable extend beyond the local block
	// local varibles are created in stack but in this case they are created in heap
	*ptr = &x
	fmt.Println("Address of x: ", &x)
}
