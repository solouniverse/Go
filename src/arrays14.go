package main

import "fmt"

func main() {
	var arr = [3]int{1, 2, 3}
	fmt.Println("Main array", arr)
	callByVal(arr)
	fmt.Println("Main array", arr)
}

func callByVal(arr [3]int) {
	fmt.Println("CallByVal", arr)
	arr[1] = 45
	fmt.Println("CallByVal", arr)
}
