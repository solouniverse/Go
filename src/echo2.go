package main

import (
	"os"
	"strings"
)
import "fmt"

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
