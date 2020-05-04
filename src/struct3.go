package main

import "fmt"

type student struct {
	name  string
	age   int
	class int
	perc  float64
}

func main() {
	var s1 student
	s1.name = "Om"
	s1.age = 19
	s1.class = 10
	s1.perc = 35.5

	fmt.Printf("%v\n", s1)
	fmt.Printf("Datatype of student: %T\n", s1)
}
