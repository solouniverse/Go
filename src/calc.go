package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.000
	b := 3

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))

	c := a + float64(b)

	fmt.Println("\nC has value:", c, "and is of type", reflect.TypeOf(c))
}
