package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	secs := time.Since(start).Seconds()
	fmt.Sprintf("%.2fs", secs)
}
