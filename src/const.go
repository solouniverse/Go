package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	aa = iota
	ab
	ac
	ad
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(aa, ab, ac, ad)
}
