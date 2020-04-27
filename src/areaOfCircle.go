package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const pi = 3.14159

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter radius: ")
	for input.Scan() {
		radius, err := strconv.ParseFloat(input.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "areaOfCircle: %v\n", err)
			continue
		} else {
			fmt.Printf("Area of a circle: %g\n", pi*radius*radius)
		}
	}
}
