package main

import "fmt"

func main() {
	var arr = [...]int{0:1, 1:2, 2:3}
	fmt.Printf("Array: %v\t%T\n", arr, arr)

	arr = [...]int{0:1, 2:3}
	fmt.Printf("Array: %v\t%T\n", arr, arr)
}
