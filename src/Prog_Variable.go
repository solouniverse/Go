package main

import (
	"fmt"
	"reflect"
)

var (
	name   string
	age    int
	gender string

	name2, age2, gender2 = "solo", "19", "Male"

	name3 = "universe"
	age3 = 21
	gender3 = "Male"
)

func main() {
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Gender is set to", gender, "and is of type", reflect.TypeOf(gender))
	fmt.Println("Age is set to", age, "and is of type", reflect.TypeOf(age))
	fmt.Println("Name is set to", name2, "and is of type", reflect.TypeOf(name2))
        fmt.Println("Gender is set to", gender2, "and is of type", reflect.TypeOf(gender2))
        fmt.Println("Age is set to", age2, "and is of type", reflect.TypeOf(age2))
        fmt.Println("Name is set to", name3, "and is of type", reflect.TypeOf(name3))
        fmt.Println("Gender is set to", gender3, "and is of type", reflect.TypeOf(gender3))
        fmt.Println("Age is set to", age3, "and is of type", reflect.TypeOf(age3))


}
