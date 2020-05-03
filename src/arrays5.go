package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v, Length: %v, Capacity: %v\n", arr, len(arr), cap(arr))
}
