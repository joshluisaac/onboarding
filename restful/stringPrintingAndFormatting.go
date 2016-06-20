package main

import "fmt"

type Payload struct {
	PayloadData Data
}

type Data map[int]int

func main() {

	// Sprintf formats according to a format specifier and returns the resulting string
	concatenated := fmt.Sprintf("a %s", "string")
	fmt.Println(concatenated)
	payload_data := map[int]int{1: 2}
	d := Payload{payload_data}
	fmt.Println(d)
	fmt.Printf("%+v\n", d) // when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%#v\n", d) // a Go-syntax representation of the value
	fmt.Printf("%T\n", d)  // a Go-syntax representation of the type of the value

	//fmt.Printf("%v\n", string(response))

}
