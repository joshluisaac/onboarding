package pointer

import (
	"fmt"
	"unsafe"
)
// pointers are variables that store the address of another variable
// updateValue takes a memory address updateValue(&whatEver)
// *someVal is an integer pointer, which means it points to something else
// *someVal is an integer pointer to val, using *someVal we can modify/update the value of val
func updateValue200(someVal *int) {
	fmt.Println("Value of the pointer:", someVal)
	fmt.Println("Dereferenced pointer variable:", *someVal)
	fmt.Println("Memory address of pointer variable:", &someVal) //memory address of pointer variable is separate from the address of the variable it is pointing to
	*someVal = *someVal + 100
	fmt.Println("Final value of the pointer after modification:", someVal)
	fmt.Println("Final dereferenced pointer variable after modification:", *someVal)
	fmt.Println("Final memory address of pointer variable after modification:", &someVal)
}

func Pointer200() {
	//val := 1000
	var val int
	val = 1000

	fmt.Printf("Size of val: %d bytes\n", unsafe.Sizeof(val))
	fmt.Println("Initial value of val:", val)
	fmt.Println("Initial address of val:", &val)
	fmt.Println("++++++++Start of function call++++++++")
	// pass in the address of the val variable
	// val passed by reference, this means that the original copy of "val" can be de-referenced and modified using an integer pointer
	// updateValue is called/executed or implemented by reference, as in, updateValue(&whatEver)
	updateValue200(&val)
	fmt.Println("++++++++End of function call++++++++")
	fmt.Println("Final value of val:", val)
	fmt.Println("Final address of val:", &val)
}
