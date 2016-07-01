package main

// Golang : Convert []string to []byte examples
// You need a fast way to convert a []string to []byte type.
// To use in situations such as storing text data into a random access file or other type of
// data manipulation that requires the input data to be in []byte type.

import (
	"fmt"
	"strings"
)

func main() {

	stringSlice := []string{"hello", "bye"}

	stringByte := "\x00" + strings.Join(stringSlice, "\x20\x00") // x20 = space and x00 = null

	fmt.Println([]byte(stringByte))

	fmt.Println(string([]byte(stringByte)))
}
