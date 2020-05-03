package main

import "fmt"

func main() {
	var arr = [4][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1}}
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Type: %T\n", arr)
}
