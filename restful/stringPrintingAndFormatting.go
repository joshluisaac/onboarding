package main

import "fmt"

type payload struct {
	payloadData data
}

type data map[int]int

func main() {

	// Sprintf formats according to a format specifier and returns the resulting string
	concatenated := fmt.Sprintf("a %s", "string")
	fmt.Println(concatenated)
	payloadData := map[int]int{1: 2}
	d := payload{payloadData}
	fmt.Println(d)
	fmt.Printf("%+v\n", d)    // when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%#v\n", d)    // a Go-syntax representation of the value
	fmt.Printf("%T\n", d)     // a Go-syntax representation of the type of the value
	fmt.Printf("%x\n", 32766) // print as hex string
	//fmt.Printf("%v\n", string(response))

}
