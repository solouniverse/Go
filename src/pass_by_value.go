package main

import (
	"fmt"
)

func main() {
	lang := "Go"
	version := 1.0

	fmt.Println("Lang is", lang, "and version is", version)

	changeVersion(version)

	fmt.Println("\nVersion is unchanged in calling function", version)
}

func changeVersion(version float64) {
	version = 1.6

	fmt.Println("\nTrying to change the version in called function", version)

	return version
}
