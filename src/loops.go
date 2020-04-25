package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("/****** Loop1 ******/")
	loop1()
	fmt.Println("/****** Loop2 ******/")
	loop2()
	fmt.Println("/****** Loop3 ******/")
	loop3()
}

func loop1() {
	var n = 5
	sum := 0

	fmt.Println("index, sum")
	for i := 1; i <= n; i++ {
		sum = sum + i
		fmt.Println(i, sum)
	}

	fmt.Println("TotalSum", sum)
}

func loop2() {
	result := ""

	// i := 0 command line argument represent the executable build from tmp folder.
	for i := 1; i < len(os.Args); i++ {
		result += os.Args[i] + " "
	}

	if len(os.Args) > 1 {
		fmt.Println(result)
	} else {
		fmt.Println("Zero command line arguments")
	}
}

func loop3() {
	var n = 10

	for i := 1; i <= n; i+=2 {
		fmt.Println(i)
	}
}
