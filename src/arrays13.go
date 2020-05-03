package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [4]int{1, 2, 3, 4}

	fmt.Printf("arr1 == arr2: %v\n", arr1 == arr2)
}
