package main

import "fmt"

func updateValue(val int) {
	fmt.Println("Initial value of val in fxn updateValue:", val)
	fmt.Println("Initial address of val in fxn updateValue:", &val)
	val = val + 100
	fmt.Println("Final value of val in fxn updateValue after modification:", val)
	fmt.Println("Final address of val in fxn updateValue after modification:", &val)
}

// 1100 may have been expected
// but by default when you pass an argument to a function in Go it will be copied by value.
// This simply means a “copy” of the val variable is used by the updateValue function.
func main() {
	val := 1000
	fmt.Println("Initial value of val:", val)
	fmt.Println("Initial address of val:", &val)
	fmt.Println("++++++++Start of function call++++++++")
	updateValue(val)
	fmt.Println("++++++++End of function call++++++++")
	fmt.Println("Final value of val:", val)
	fmt.Println("Final address of val:", &val)
}
