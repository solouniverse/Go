package main

import "fmt"

func main() {
	fmt.Println("Factorial")

	n := 5
	fact := 1

	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println(fact)
}
