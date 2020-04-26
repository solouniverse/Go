package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i+=1 {
			s = s + sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	} else {
		fmt.Println("No command line arguments passed")
	}
}
