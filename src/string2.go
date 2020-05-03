package main

import "fmt"

func main() {
	line := "This is a Go programming language"

	fmt.Println("11th index is", line[10])

	fmt.Printf("11th index is %c\n", line[10])

	fmt.Printf("length of string is %d\n", len(line))
}
