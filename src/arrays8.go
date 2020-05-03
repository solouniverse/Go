package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3}
	fmt.Printf("Array: %v\n", arr)

	arr = [3]int{10, 11}
	fmt.Printf("Array: %v\n", arr)
}
