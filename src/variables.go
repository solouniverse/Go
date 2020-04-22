package main

import (
	"fmt"
	"reflect"
)

// declaring at package level
var (
	firstName string
	lastName, gender string
	height float64
	weight float64
	
	// can be initialized in multiple lines
	firstName2, lastName2, gender2, height2, weight2 = "Hare Krishna", "Hare Rama", "unisex", 99.99999, 99.999
	
	
)

func main() {
	fmt.Println("First Name: ", firstName)
	fmt.Println("Last Name: ", lastName, "and is of type", reflect.TypeOf(lastName))
	fmt.Println("Gender: ", gender)
	fmt.Println("Height: ", height, reflect.TypeOf(height))
	fmt.Println("Weight: ", weight)
	
	
	fmt.Println("First Name: ", firstName2)
	fmt.Println("Last Name: ", lastName2, "and is of type", reflect.TypeOf(lastName2))
	fmt.Println("Gender: ", gender2)
	fmt.Println("Height: ", height2, reflect.TypeOf(height2))
	fmt.Println("Weight: ", weight2)
	
	// variable at function level
	a := 10.123
	b := 2
	
	fmt.Println("\n variable a is type", reflect.TypeOf(a), "and b is type", reflect.TypeOf(b))
	
	//c := a + b //invalid operation: a + b (mismatched types float64 and int)
	c := int(a) + b //cast to integer
	
	fmt.Println("\n c has value:", c, "and is of type", reflect.TypeOf(c))
}
