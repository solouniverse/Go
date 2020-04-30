package main

import "fmt"

func main() {
	x := 1
	p := &x

	fmt.Println("Value of pointer p: ", *p)

	*p = 2

	fmt.Println("Value of variable x:", x)
}
