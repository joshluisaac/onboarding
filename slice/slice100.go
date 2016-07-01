package main

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
