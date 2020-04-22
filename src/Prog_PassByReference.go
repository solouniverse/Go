package main

import "fmt"

func main() {
	lang := "Go"
	version := 1.5
	ptr := &version

	//fmt.Printf("Lang:%v", lang, "Version:%v", version, "Type:%T", version)

	fmt.Println("Lang:", lang, "Version:", version)

	// By reference ampersand &
	changeVersion(ptr)

	fmt.Println("\nVersion changed in calling function by passing reference.", version)
}

func changeVersion(version *float64) float64 {
	// By dereference asterisk *
	*version = 1.6

	fmt.Println("\nVersion changed in called function by drefernce.", *version)

	return *version
}
