package main

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
