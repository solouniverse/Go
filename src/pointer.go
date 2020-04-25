package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Sri"
	version := 1.6
	ptr := &version

	fmt.Println("Namaskaram", name, "and is of Type", reflect.TypeOf(name))
	fmt.Println("Go Version", version, "and is of Type", reflect.TypeOf(version))
	fmt.Println("Memory address of *name* variable is", &name)
	fmt.Println("Memory address of *version* variable is", &version)
	fmt.Println("Memory address of *version* variable is", ptr, "and value of *version* using pointer is", *ptr)
}
