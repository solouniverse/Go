package main

import "fmt"

func main() {
	var student struct {
		name string
		age int
		class int
		perc float64
	}

	student.name = "Om"
	student.age = 19
	student.class = 10
	student.perc = 35.5

	fmt.Printf("%v\n", student)
}
