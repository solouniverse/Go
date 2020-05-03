package main

import "fmt"

func main() {
	var arr = [3]int{1, 2, 3}
	fmt.Println("Main array", arr)
	callByVal(arr)
	fmt.Println("Main array", arr)
	callByRef(&arr)
	fmt.Println("Main array", arr)
}

func callByVal(arr [3]int) {
	fmt.Println("CallByVal", arr)
	arr[1] = 45
	fmt.Println("CallByVal", arr)
}

func callByRef(arr *[3]int) {
	fmt.Println("CallByRef", *arr)
	(*arr)[0] = 897
	fmt.Println("CallByRef", *arr)
}
