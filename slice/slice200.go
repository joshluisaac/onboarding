package main

// Golang : Convert []string to []byte examples
// You need a fast way to convert a []string to []byte type.
// To use in situations such as storing text data into a random access file or other type of
// data manipulation that requires the input data to be in []byte type.

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {

	stringSlice := []string{"hello", "bye"}

	buffer := &bytes.Buffer{}

	gob.NewEncoder(buffer).Encode(stringSlice)
	byteSlice := buffer.Bytes()
	fmt.Printf("%q\n", byteSlice)

	fmt.Println("---------------------------")

	backToStringSlice := []string{}
	gob.NewDecoder(buffer).Decode(&backToStringSlice)
	fmt.Printf("%v\n", backToStringSlice)
}
